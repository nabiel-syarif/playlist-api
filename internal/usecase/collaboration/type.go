package collaboration

import (
	"context"

	collaborationModel "github.com/nabiel-syarif/playlist-api/internal/model/collaboration"
	userRepo "github.com/nabiel-syarif/playlist-api/internal/repo/auth"
	collaborationRepo "github.com/nabiel-syarif/playlist-api/internal/repo/collaboration"
	playlsitRepo "github.com/nabiel-syarif/playlist-api/internal/repo/playlist"
	"github.com/nabiel-syarif/playlist-api/pkg/db"
	workerpool "github.com/nabiel-syarif/playlist-api/pkg/worker-pool"
)

//go:generate moq -out collaboration_mock_test.go . Repo
type Repo interface {
	AddCollaborator(ctx context.Context, playlistOwnerId, playlistId int, req []collaborationModel.AddCollaboratorRequest) error
	RemoveCollaborator(ctx context.Context, playlistId int, req []collaborationModel.RemoveCollaboratorRequest) error
}

type Usecase interface {
	AddCollaborator(ctx context.Context, playlistId, senderUserId int, req []collaborationModel.AddCollaboratorRequest) error
	RemoveCollaborator(ctx context.Context, playlistId, userId int, req []collaborationModel.RemoveCollaboratorRequest) error
}

type usecase struct {
	playlistRepo playlsitRepo.Repo
	collabRepo   collaborationRepo.Repo
	userRepo     userRepo.Repo
	wp           *workerpool.WorkerPool
	db           db.DB
}

func New(db db.DB, pRepo playlsitRepo.Repo, cRepo collaborationRepo.Repo, uRepo userRepo.Repo, wp *workerpool.WorkerPool) Usecase {
	return &usecase{
		playlistRepo: pRepo,
		collabRepo:   cRepo,
		userRepo:     uRepo,
		wp:           wp,
		db:           db,
	}
}
