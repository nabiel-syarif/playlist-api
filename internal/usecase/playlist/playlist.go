package playlist

import (
	"context"

	"github.com/jackc/pgx/v4"
	playlistModel "github.com/nabiel-syarif/playlist-api/internal/model/playlist"
	songModel "github.com/nabiel-syarif/playlist-api/internal/model/song"
	playlistRepo "github.com/nabiel-syarif/playlist-api/internal/repo/playlist"
	playlistErr "github.com/nabiel-syarif/playlist-api/pkg/error/playlist"
)

//go:generate moq -out playlist_mock_test.go . Repo
type Repo interface {
	SavePlaylist(ctx context.Context, p playlistModel.AddPlaylistRequest) (playlistModel.PlaylistAggregated, error)
	ListPlaylists(ctx context.Context, userId int) ([]playlistModel.PlaylistAggregated, error)
	GetPlaylistById(ctx context.Context, playlistId int) (playlistModel.Playlist, error)
	GetSongsFromPlaylistId(ctx context.Context, playlistId int) ([]songModel.Song, error)
	UpdatePlaylist(ctx context.Context, p playlistModel.UpdatePlaylistRequest) (playlistModel.Playlist, error)
	DeletePlaylist(ctx context.Context, playlistId int) error
	AddSongToPlaylist(ctx context.Context, playlistId int, songId int) error
	RemoveSongFromPlaylist(ctx context.Context, playlistId int, songId int) error
	CheckPlaylistAccess(ctx context.Context, playlistId, userid int) error
}

var _ Usecase = &usecase{}

func New(repo playlistRepo.Repo) Usecase {
	return &usecase{
		repo: repo,
	}
}

func (uc *usecase) SavePlaylist(ctx context.Context, req playlistModel.AddPlaylistRequest) (playlistModel.PlaylistAggregated, error) {
	return uc.repo.SavePlaylist(ctx, req)
}

func (uc *usecase) ListPlaylists(ctx context.Context, userId int) ([]playlistModel.PlaylistAggregated, error) {
	playlists, err := uc.repo.ListPlaylists(ctx, userId)
	if err != nil {
		return nil, err
	}

	return playlists, nil
}

func (uc *usecase) GetPlaylistById(ctx context.Context, userId, playlistId int) (playlistModel.PlaylistAggregated, error) {
	if err := uc.CheckPlaylistAccess(ctx, userId, playlistId); err != nil {
		if err == pgx.ErrNoRows {
			return playlistModel.PlaylistAggregated{}, playlistErr.ErrPlaylistNotFound
		}
		return playlistModel.PlaylistAggregated{}, err
	}

	var playlist playlistModel.PlaylistAggregated
	playlistModel, err := uc.repo.GetPlaylistById(ctx, playlistId)
	if err != nil {
		if err == pgx.ErrNoRows {
			return playlist, playlistErr.ErrPlaylistNotFound
		}
		return playlist, err
	}

	playlist.Id = playlistModel.Id
	playlist.Name = playlistModel.Name
	playlist.CreatedAt = playlistModel.CreatedAt
	playlist.UpdatedAt = playlistModel.UpdatedAt
	playlist.Songs, err = uc.repo.GetSongsFromPlaylistId(ctx, playlistId)
	return playlist, err
}

func (uc *usecase) UpdatePlaylist(ctx context.Context, req playlistModel.UpdatePlaylistRequest) (playlistModel.PlaylistAggregated, error) {
	if err := uc.CheckPlaylistAccess(ctx, req.Owner, req.Id); err != nil {
		return playlistModel.PlaylistAggregated{}, err
	}

	var playlist playlistModel.PlaylistAggregated

	playlistModel, err := uc.repo.UpdatePlaylist(ctx, req)
	if err != nil {
		if err == pgx.ErrNoRows {
			return playlist, playlistErr.ErrPlaylistNotFound
		}
		return playlist, err
	}

	playlist.Id = playlistModel.Id
	playlist.Name = playlistModel.Name
	playlist.CreatedAt = playlistModel.CreatedAt
	playlist.UpdatedAt = playlistModel.UpdatedAt

	songs, err := uc.repo.GetSongsFromPlaylistId(ctx, req.Id)
	if err != nil {
		return playlist, err
	}
	playlist.Songs = songs

	return playlist, nil
}

func (uc *usecase) DeletePlaylist(ctx context.Context, userId, playlistId int) error {
	if err := uc.CheckPlaylistAccess(ctx, userId, playlistId); err != nil {
		return err
	}

	return uc.repo.DeletePlaylist(ctx, playlistId)
}

func (uc *usecase) CheckPlaylistAccess(ctx context.Context, userId, playlistId int) error {
	err := uc.repo.CheckPlaylistAccess(ctx, playlistId, userId)
	if err != nil {
		return err
	}

	return nil
}

func (uc *usecase) AttachSongToPlaylist(ctx context.Context, userId, playlistId, songId int) error {
	if err := uc.CheckPlaylistAccess(ctx, userId, playlistId); err != nil {
		return err
	}

	return uc.repo.AddSongToPlaylist(ctx, playlistId, songId)
}

func (uc *usecase) DetachSongFromPlaylist(ctx context.Context, userId int, playlistId int, songId int) error {
	if err := uc.CheckPlaylistAccess(ctx, userId, playlistId); err != nil {
		return err
	}

	return uc.repo.RemoveSongFromPlaylist(ctx, playlistId, songId)
}
