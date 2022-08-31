package song

import (
	"context"

	modelSong "github.com/nabiel-syarif/playlist-api/internal/model/song"
	songRepo "github.com/nabiel-syarif/playlist-api/internal/repo/song"
)

type usecase struct {
	repo songRepo.Repo
}

type Usecase interface {
	GetSongs(ctx context.Context) ([]modelSong.Song, error)
	GetSongById(ctx context.Context, id int) (modelSong.Song, error)
	SaveSong(ctx context.Context, song modelSong.InsertSongRequest) (modelSong.Song, error)
	UpdateSong(ctx context.Context, songId int, song modelSong.InsertSongRequest) (modelSong.Song, error)
	DeleteSong(ctx context.Context, songId int) (error)
}
