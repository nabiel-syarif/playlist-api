package song

import (
	"context"

	modelSong "github.com/nabiel-syarif/playlist-api/internal/model/song"
	"github.com/nabiel-syarif/playlist-api/pkg/db"
)

type repo struct {
	db db.DB
}

type Repo interface {
	InsertSong(ctx context.Context, song modelSong.InsertSongRequest) (modelSong.Song, error)
	GetSongById(ctx context.Context, id int) (modelSong.Song, error)
	GetSongs(ctx context.Context) ([]modelSong.Song, error)
	UpdateSongById(ctx context.Context, id int, song modelSong.InsertSongRequest) (modelSong.Song, error)
	DeleteSongById(ctx context.Context, id int) error
}

func New(db db.DB) Repo {
	return &repo{
		db: db,
	}
}

var _ Repo = &repo{}

func (repo *repo) InsertSong(ctx context.Context, song modelSong.InsertSongRequest) (modelSong.Song, error) {
	sql := `INSERT INTO songs(title, performer, genre, duration) VALUES ($1, $2, $3, $4) RETURNING song_id`
	row := repo.db.QueryRow(ctx, sql, song.Title, song.Performer, song.Genre, song.Duration)

	var newSong modelSong.Song
	var id int
	err := row.Scan(&id)
	if err != nil {
		return newSong, err
	}

	sql = `SELECT * FROM songs WHERE song_id = $1 LIMIT 1`
	row = repo.db.QueryRow(ctx, sql, id)
	err = row.Scan(&newSong.Id, &newSong.Title, &newSong.Performer, &newSong.Genre, &newSong.Duration, &newSong.CreatedAt, &newSong.UpdatedAt)
	return newSong, err
}

func (repo *repo) GetSongById(ctx context.Context, id int) (modelSong.Song, error) {
	row := repo.db.QueryRow(ctx, "SELECT * FROM songs WHERE song_id = $1", id)
	var song modelSong.Song
	err := row.Scan(&song.Id, &song.Title, &song.Performer, &song.Genre, &song.Duration, &song.CreatedAt, &song.UpdatedAt)
	return song, err
}

func (repo *repo) GetSongs(ctx context.Context) ([]modelSong.Song, error) {
	rows, err := repo.db.Query(ctx, "SELECT * FROM songs")
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

func (repo *repo) UpdateSongById(ctx context.Context, id int, song modelSong.InsertSongRequest) (modelSong.Song, error) {
	sql := `UPDATE songs SET title = $1, performer = $2, genre = $3, duration = $4 WHERE song_id = $5`
	var updatedSong modelSong.Song
	_, err := repo.db.Exec(ctx, sql, song.Title, song.Performer, song.Genre, song.Duration, id)
	if err != nil {
		return updatedSong, err
	}
	sql = `SELECT * FROM songs WHERE song_id = $1 LIMIT 1`
	row := repo.db.QueryRow(ctx, sql, id)
	err = row.Scan(&updatedSong.Id, &updatedSong.Title, &updatedSong.Performer, &updatedSong.Genre, &updatedSong.Duration, &updatedSong.CreatedAt, &updatedSong.UpdatedAt)
	return updatedSong, err
}

func (repo *repo) DeleteSongById(ctx context.Context, id int) error {
	sql := `DELETE FROM songs WHERE song_id = $1`
	_, err := repo.db.Exec(ctx, sql, id)
	return err
}
