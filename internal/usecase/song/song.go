package song

import (
	"context"

	"github.com/jackc/pgx/v4"
	modelSong "github.com/nabiel-syarif/playlist-api/internal/model/song"
	songRepo "github.com/nabiel-syarif/playlist-api/internal/repo/song"
	songErr "github.com/nabiel-syarif/playlist-api/pkg/error/song"
)

var _ Usecase = &usecase{}

//go:generate moq -out song_mock_test.go . Repo
type Repo interface {
	InsertSong(ctx context.Context, song modelSong.InsertSongRequest) (modelSong.Song, error)
	GetSongById(ctx context.Context, id int) (modelSong.Song, error)
	GetSongs(ctx context.Context) ([]modelSong.Song, error)
	UpdateSongById(ctx context.Context, id int, song modelSong.InsertSongRequest) (modelSong.Song, error)
	DeleteSongById(ctx context.Context, id int) error
}

func New(repo songRepo.Repo) Usecase {
	return &usecase{
		repo: repo,
	}
}

func (uc *usecase) SaveSong(ctx context.Context, song modelSong.InsertSongRequest) (modelSong.Song, error) {
	newSong, err := uc.repo.InsertSong(ctx, song)
	return newSong, err
}

func (uc *usecase) GetSongs(ctx context.Context) ([]modelSong.Song, error) {
	songs, err := uc.repo.GetSongs(ctx)

	if err != nil {
		return nil, err
	}

	return songs, nil
}

func (uc *usecase) GetSongById(ctx context.Context, id int) (modelSong.Song, error) {

	song, err := uc.repo.GetSongById(ctx, id)
	if err != nil {
		if err == pgx.ErrNoRows {
			return modelSong.Song{}, songErr.ErrSongNotFound
		}

		return modelSong.Song{}, err
	}

	return song, nil
}

func (uc *usecase) UpdateSong(ctx context.Context, songId int, song modelSong.InsertSongRequest) (modelSong.Song, error) {
	return uc.repo.UpdateSongById(ctx, songId, song)
}

func (uc *usecase) DeleteSong(ctx context.Context, songId int) error {
	return uc.repo.DeleteSongById(ctx, songId)
}
