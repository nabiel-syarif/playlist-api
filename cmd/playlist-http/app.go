package main

import (
	"context"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/nabiel-syarif/playlist-api/internal/config"

	authHandler "github.com/nabiel-syarif/playlist-api/internal/handler/auth"
	authRepo "github.com/nabiel-syarif/playlist-api/internal/repo/auth"
	authUc "github.com/nabiel-syarif/playlist-api/internal/usecase/auth"

	songHandler "github.com/nabiel-syarif/playlist-api/internal/handler/song"
	songRepo "github.com/nabiel-syarif/playlist-api/internal/repo/song"
	songUc "github.com/nabiel-syarif/playlist-api/internal/usecase/song"

	playlistHandler "github.com/nabiel-syarif/playlist-api/internal/handler/playlist"
	playlistRepo "github.com/nabiel-syarif/playlist-api/internal/repo/playlist"
	playlistUc "github.com/nabiel-syarif/playlist-api/internal/usecase/playlist"

	collaborationHandler "github.com/nabiel-syarif/playlist-api/internal/handler/collaboration"
	collaborationRepo "github.com/nabiel-syarif/playlist-api/internal/repo/collaboration"
	collaborationUc "github.com/nabiel-syarif/playlist-api/internal/usecase/collaboration"

	"github.com/nabiel-syarif/playlist-api/pkg/jwt"
	"github.com/nabiel-syarif/playlist-api/pkg/mailer"
	workerpool "github.com/nabiel-syarif/playlist-api/pkg/worker-pool"
)

func startApp(cfg *config.Config) error {
	mailer.SetMailConfig(cfg.MailConfig)
	conn, err := pgxpool.Connect(context.Background(), cfg.Database.PostgreURL)
	if err != nil {
		return err
	}
	defer conn.Close()

	jwtHelper := jwt.JwtHelper{
		Config: struct {
			SecretKey       string
			TokenExpiration int
		}{
			SecretKey:       cfg.JwtConfig.SecretKey,
			TokenExpiration: cfg.JwtConfig.ExpirationTime,
		},
	}
	authRepo := authRepo.New(conn)
	authUc := authUc.New(authRepo, jwtHelper)
	authHandler := authHandler.New(authUc)

	songRepo := songRepo.New(conn)
	songUc := songUc.New(songRepo)
	songHandler := songHandler.New(songUc)

	playlistRepo := playlistRepo.New(conn)
	playlistUc := playlistUc.New(playlistRepo)
	playlistHandler := playlistHandler.New(playlistUc)

	wp := workerpool.New(cfg.WorkerPoolConfig.WorkerCounts)
	ctxWp, cancelWp := context.WithCancel(context.Background())
	wp.Run(ctxWp)

	collabRepo := collaborationRepo.New(conn)
	collabUc := collaborationUc.New(conn, playlistRepo, collabRepo, authRepo, wp)
	collabHandler := collaborationHandler.New(collabUc)

	router := newRoute(Handlers{
		AuthHandler:     authHandler,
		SongHandler:     songHandler,
		PlaylistHandler: playlistHandler,
		CollabHandler:   collabHandler,

		JwtHelper: jwtHelper,
	})

	return startServer(router, cfg, []OnShutdownCallback{
		func() error {
			if wp.IsWorkerPoolWaitingQueueEmpty() {
				cancelWp()
			}
			wp.Wait()
			wp.Close()
			return nil
		},
	})
}
