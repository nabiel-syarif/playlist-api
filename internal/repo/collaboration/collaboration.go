package collaboration

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4"
	collaborationModel "github.com/nabiel-syarif/playlist-api/internal/model/collaboration"
	"github.com/nabiel-syarif/playlist-api/pkg/db"
)

type repo struct {
	db db.DB
}

type Repo interface {
	AddCollaborator(ctx context.Context, playlistOwnerId, playlistId int, req []collaborationModel.AddCollaboratorRequest) error
	RemoveCollaborator(ctx context.Context, playlistId int, req []collaborationModel.RemoveCollaboratorRequest) error
}

var _ Repo = &repo{}

func New(db db.DB) Repo {
	return &repo{db: db}
}

func (r *repo) AddCollaborator(ctx context.Context, playlistOwnerId, playlistId int, req []collaborationModel.AddCollaboratorRequest) error {
	sql := `INSERT INTO playlists_collaborations(playlist_id, user_id) VALUES `
	params := []interface{}{}
	for i, v := range req {
		if v.UserId == playlistOwnerId {
			continue
		}
		if i == len(req)-1 {
			sql += fmt.Sprintf("($%d, $%d)", (i*2 + 1), (i*2 + 2))
		} else {
			sql += fmt.Sprintf("($%d, $%d), ", (i*2 + 1), (i*2 + 2))
		}
		params = append(params, playlistId, v.UserId)
	}

	if len(params) == 0 {
		return pgx.ErrNoRows
	}

	_, err := r.db.Exec(ctx, sql, params...)
	return err
}

func (r *repo) RemoveCollaborator(ctx context.Context, playlistId int, req []collaborationModel.RemoveCollaboratorRequest) error {
	sql := `DELETE FROM playlists_collaborations WHERE playlist_id = $1 AND user_id IN (`
	params := make([]interface{}, 0)
	params = append(params, playlistId)
	for i, v := range req {
		if i == len(req)-1 {
			sql += fmt.Sprintf("$%d", i+2)
		} else {
			sql += fmt.Sprintf("$%d, ", i+2)
		}
		params = append(params, v.UserId)
	}
	sql += `)`

	if len(params) == 1 {
		return pgx.ErrNoRows
	}

	_, err := r.db.Exec(ctx, sql, params...)
	return err
}
