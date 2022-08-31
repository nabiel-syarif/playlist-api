package main

import (
	"github.com/go-chi/chi/v5"
	authHandler "github.com/nabiel-syarif/playlist-api/internal/handler/auth"
	collabHandler "github.com/nabiel-syarif/playlist-api/internal/handler/collaboration"
	"github.com/nabiel-syarif/playlist-api/internal/handler/middlewares"
	playlistHandler "github.com/nabiel-syarif/playlist-api/internal/handler/playlist"
	songHandler "github.com/nabiel-syarif/playlist-api/internal/handler/song"
	"github.com/nabiel-syarif/playlist-api/pkg/jwt"
)

type Handlers struct {
	AuthHandler     authHandler.Handler
	SongHandler     songHandler.Handler
	CollabHandler   collabHandler.Handler
	PlaylistHandler playlistHandler.Handler
	JwtHelper       jwt.JwtHelper
}

func newRoute(handlers Handlers) *chi.Mux {
	router := chi.NewRouter()

	router.Group(authRoutes(handlers.AuthHandler))
	// protected routes
	router.Group(protectedRoutes(handlers))

	return router
}

func authRoutes(handler authHandler.Handler) func(r chi.Router) {
	return func(r chi.Router) {
		r.Post("/v1/auth/register", handler.Register)
		r.Post("/v1/auth/login", handler.Login)
	}
}

func protectedRoutes(handlers Handlers) func(router chi.Router) {
	return func(router chi.Router) {
		router.Use(middlewares.AuthOnly(handlers.JwtHelper))
		router.Group(songRoutes(handlers.SongHandler))
		router.Group(playlistRoutes(handlers.PlaylistHandler))
		router.Group(collaborationRoutes(handlers.CollabHandler))
	}
}

func songRoutes(handler songHandler.Handler) func(r chi.Router) {
	return func(r chi.Router) {
		r.Get("/v1/songs", handler.ListSong)
		r.Post("/v1/songs", handler.SaveSong)
		r.Get("/v1/songs/{songId}", handler.GetSongById)

		r.Put("/v1/songs/{songId}", handler.UpdateSong)
		r.Delete("/v1/songs/{songId}", handler.DeleteSong)
	}
}

func playlistRoutes(handler playlistHandler.Handler) func(r chi.Router) {
	return func(r chi.Router) {
		r.Post("/v1/playlists", handler.SavePlaylist)
		r.Get("/v1/playlists", handler.ListPlaylist)

		r.Post("/v1/playlists/attach-songs", handler.AttachSongToPlaylist)
		r.Delete("/v1/playlists/detach-songs", handler.DetachSongFromPlaylist)

		r.Get("/v1/playlists/{playlistId}", handler.GetPlaylistById)
		r.Put("/v1/playlists/{playlistId}", handler.UpdatePlaylist)
		r.Delete("/v1/playlists/{playlistId}", handler.DeletePlaylist)
	}
}

func collaborationRoutes(handler collabHandler.Handler) func(r chi.Router) {
	return func(r chi.Router) {
		r.Post("/v1/playlists/{playlistId}/add-collaborators", handler.AddCollaborator)
		r.Post("/v1/playlists/{playlistId}/remove-collaborators", handler.RemoveCollaborator)
	}
}
