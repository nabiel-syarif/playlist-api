package playlist

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4"
	modelPlaylist "github.com/nabiel-syarif/playlist-api/internal/model/playlist"
	modelSong "github.com/nabiel-syarif/playlist-api/internal/model/song"
	db "github.com/nabiel-syarif/playlist-api/pkg/db"
	playlistErr "github.com/nabiel-syarif/playlist-api/pkg/error/playlist"
)

type Repo interface {
	SavePlaylist(ctx context.Context, p modelPlaylist.AddPlaylistRequest) (modelPlaylist.PlaylistAggregated, error)
	ListPlaylists(ctx context.Context, userId int) ([]modelPlaylist.PlaylistAggregated, error)
	GetPlaylistById(ctx context.Context, playlistId int) (modelPlaylist.Playlist, error)
	GetSongsFromPlaylistId(ctx context.Context, playlistId int) ([]modelSong.Song, error)
	UpdatePlaylist(ctx context.Context, p modelPlaylist.UpdatePlaylistRequest) (modelPlaylist.Playlist, error)
	DeletePlaylist(ctx context.Context, playlistId int) error
	AddSongToPlaylist(ctx context.Context, playlistId int, songId int) error
	RemoveSongFromPlaylist(ctx context.Context, playlistId int, songId int) error
	CheckPlaylistAccess(ctx context.Context, playlistId, userid int) error
}

type repo struct {
	db db.DB
}

var _ Repo = &repo{}

func New(db db.DB) Repo {
	return &repo{
		db: db,
	}
}

func (repo *repo) SavePlaylist(ctx context.Context, p modelPlaylist.AddPlaylistRequest) (modelPlaylist.PlaylistAggregated, error) {
	sql := `INSERT INTO playlists(name, owner_id) VALUES ($1, $2) RETURNING playlist_id`
	row := repo.db.QueryRow(ctx, sql, p.Name, p.Owner)

	var id int
	var playlist modelPlaylist.PlaylistAggregated
	err := row.Scan(&id)

	if err != nil {
		return playlist, err
	}

	sql = `SELECT playlist_id, name, created_at, updated_at FROM playlists WHERE playlist_id = $1`
	row = repo.db.QueryRow(ctx, sql, id)
	err = row.Scan(&playlist.Id, &playlist.Name, &playlist.CreatedAt, &playlist.UpdatedAt)
	if err != nil {
		return playlist, err
	}
	playlist.Songs = make([]modelSong.Song, 0)
	return playlist, nil
}

func (repo *repo) ListPlaylists(ctx context.Context, userId int) ([]modelPlaylist.PlaylistAggregated, error) {
	sql := `SELECT p.playlist_id, p.name, p.created_at, p.updated_at FROM playlists p LEFT JOIN playlists_collaborations pc ON pc.playlist_id = p.playlist_id WHERE p.owner_id = $1 OR pc.user_id = $1`
	rows, err := repo.db.Query(ctx, sql, userId)
	if err != nil {
		return nil, err
	}

	playlistsMap := make(map[int]*modelPlaylist.PlaylistAggregated)
	playlistIds := make([]interface{}, 0)
	for rows.Next() {
		var playlist modelPlaylist.PlaylistAggregated
		err := rows.Scan(&playlist.Id, &playlist.Name, &playlist.CreatedAt, &playlist.UpdatedAt)
		if err != nil {
			return nil, err
		}

		playlist.Songs = make([]modelSong.Song, 0)
		playlistsMap[playlist.Id] = &playlist
		playlistIds = append(playlistIds, playlist.Id)
	}

	if len(playlistIds) == 0 {
		return nil, pgx.ErrNoRows
	}

	sql = `SELECT s.*, ps.playlist_id FROM playlists_songs ps INNER JOIN songs s ON s.song_id = ps.song_id WHERE ps.playlist_id IN (`
	for i := range playlistIds {
		sql += fmt.Sprintf("$%d ", i+1)
	}
	sql += `)`

	rows, err = repo.db.Query(ctx, sql, playlistIds...)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var playlistId int
		var song modelSong.Song
		err := rows.Scan(&song.Id, &song.Title, &song.Performer, &song.Genre, &song.Duration, &song.CreatedAt, &song.UpdatedAt, &playlistId)
		if err != nil {
			return nil, err
		}
		playlist, ok := playlistsMap[playlistId]
		if ok {
			playlist.Songs = append(playlist.Songs, song)
		}
	}

	playlists := make([]modelPlaylist.PlaylistAggregated, 0)
	for _, v := range playlistsMap {
		playlists = append(playlists, *v)
	}

	return playlists, nil
}

func (repo *repo) GetPlaylistById(ctx context.Context, playlistId int) (modelPlaylist.Playlist, error) {
	sql := `SELECT * FROM playlists WHERE playlist_id = $1 LIMIT 1`
	row := repo.db.QueryRow(ctx, sql, playlistId)
	var playlist modelPlaylist.Playlist
	err := row.Scan(&playlist.Id, &playlist.Name, &playlist.OwnerId, &playlist.CreatedAt, &playlist.UpdatedAt)
	return playlist, err
}

func (repo *repo) UpdatePlaylist(ctx context.Context, p modelPlaylist.UpdatePlaylistRequest) (modelPlaylist.Playlist, error) {
	sql := `UPDATE playlists SET name = $1 WHERE playlist_id = $2`

	var playlist modelPlaylist.Playlist

	_, err := repo.db.Exec(ctx, sql, p.Name, p.Id)
	if err != nil {
		return playlist, err
	}

	sql = `SELECT * FROM playlists WHERE playlist_id = $1 LIMIT 1`
	row := repo.db.QueryRow(ctx, sql, p.Id)
	err = row.Scan(&playlist.Id, &playlist.Name, &playlist.OwnerId, &playlist.CreatedAt, &playlist.UpdatedAt)
	return playlist, err
}

func (repo *repo) DeletePlaylist(ctx context.Context, playlistId int) error {
	sql := `DELETE FROM playlists WHERE playlist_id = $1`
	_, err := repo.db.Exec(ctx, sql, playlistId)
	return err
}

func (repo *repo) GetSongsFromPlaylistId(ctx context.Context, playlistId int) ([]modelSong.Song, error) {
	sql := `SELECT s.* FROM playlists_songs ps INNER JOIN songs s ON s.song_id = ps.song_id WHERE ps.playlist_id = $1`
	rows, err := repo.db.Query(ctx, sql, playlistId)
	if err != nil {
		return nil, err
	}

	songs := make([]modelSong.Song, 0)
	for rows.Next() {
		var song modelSong.Song
		err := rows.Scan(&song.Id, &song.Title, &song.Performer, &song.Genre, &song.Duration, &song.CreatedAt, &song.UpdatedAt)
		if err != nil {
			return nil, err
		}
		songs = append(songs, song)
	}

	return songs, nil
}

func (repo *repo) AddSongToPlaylist(ctx context.Context, playlistId int, songId int) error {
	sql := `INSERT INTO playlists_songs(playlist_id, song_id) VALUES ($1, $2)`
	_, err := repo.db.Exec(ctx, sql, playlistId, songId)
	return err
}

func (repo *repo) RemoveSongFromPlaylist(ctx context.Context, playlistId int, songId int) error {
	sql := `DELETE FROM playlists_songs WHERE playlist_id = $1 AND song_id = $2`
	_, err := repo.db.Exec(ctx, sql, playlistId, songId)
	return err
}

func (repo *repo) CheckPlaylistAccess(ctx context.Context, playlistId, userId int) error {
	sql := `SELECT * FROM playlists p LEFT JOIN playlists_collaborations pc ON pc.playlist_id = $1 WHERE p.owner_id = $2 OR pc.user_id = $2 LIMIT 1`
	tag, err := repo.db.Exec(ctx, sql, playlistId, userId)
	if err != nil {
		return err
	}
	if tag.RowsAffected() == 0 {
		return playlistErr.ErrNotPlaylistOwner
	}

	return nil
}
