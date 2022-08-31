package playlist

import (
	"context"

	playlistModel "github.com/nabiel-syarif/playlist-api/internal/model/playlist"
	playlistRepo "github.com/nabiel-syarif/playlist-api/internal/repo/playlist"
)

type usecase struct {
	repo playlistRepo.Repo
}

type Usecase interface {
	SavePlaylist(ctx context.Context, req playlistModel.AddPlaylistRequest) (playlistModel.PlaylistAggregated, error)
	AttachSongToPlaylist(ctx context.Context, userId, playlistId, songId int) (error)
	DetachSongFromPlaylist(ctx context.Context, userId, playlistId, songId int) (error)
	ListPlaylists(ctx context.Context, userId int) ([]playlistModel.PlaylistAggregated, error)
	GetPlaylistById(ctx context.Context, userId, playlistId int) (playlistModel.PlaylistAggregated, error)
	UpdatePlaylist(ctx context.Context, req playlistModel.UpdatePlaylistRequest) (playlistModel.PlaylistAggregated, error)
	DeletePlaylist(ctx context.Context, userId, playlistId int) (error)
}
