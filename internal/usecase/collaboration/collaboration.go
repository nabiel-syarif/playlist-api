package collaboration

import (
	"context"
	"fmt"
	"log"

	collaborationModel "github.com/nabiel-syarif/playlist-api/internal/model/collaboration"
	collabErr "github.com/nabiel-syarif/playlist-api/pkg/error/collaboration"
	"github.com/nabiel-syarif/playlist-api/pkg/mailer"
	workerpool "github.com/nabiel-syarif/playlist-api/pkg/worker-pool"
)

var _ Usecase = &usecase{}

func (uc *usecase) AddCollaborator(ctx context.Context, playlistId, senderUserId int, req []collaborationModel.AddCollaboratorRequest) error {
	playlist, err := uc.playlistRepo.GetPlaylistById(ctx, playlistId)
	if err != nil {
		return err
	}

	if playlist.OwnerId != senderUserId {
		return collabErr.ErrAddCollaboratorNotPlaylistOwner
	}

	err = uc.collabRepo.AddCollaborator(ctx, playlist.OwnerId, playlistId, req)
	if err != nil {
		return err
	}

	sender, err := uc.userRepo.GetUserByUserId(ctx, senderUserId)
	if err != nil {
		return err
	}

	jobs := make([]workerpool.Job, len(req))
	for i, v := range req {
		if v.UserId == playlist.OwnerId {
			continue
		}

		jobs[i] = workerpool.Job{
			Id: i,
			Fn: func(ctx context.Context, args interface{}) (interface{}, error) {
				sliceArgs := args.([]interface{})

				jobId := sliceArgs[1]
				req := sliceArgs[0].(collaborationModel.AddCollaboratorRequest)

				log.Printf("Job %d executed\n", jobId)

				to, err := uc.userRepo.GetUserByUserId(ctx, req.UserId)
				if err != nil {
					return nil, err
				}

				log.Printf("start sending email notification to %s\n", to.Email)
				err = mailer.SendAddedAsCollaboratorEmailNotification(playlist.Name, sender.Email, sender.Name, to.Email, to.Name)
				if err != nil {
					return nil, err
				}

				return fmt.Sprintf("Sent collaborator email notification to %s\n", to.Name), nil
			},
			Args: []interface{}{v, i},
		}
	}

	go uc.wp.FromJobs(jobs)
	go func() {
		results := uc.wp.Results()
		for i := 0; i < len(jobs); i++ {
			res := <-results
			if res.Error != nil {
				log.Printf("Job %d Return error, err : %v\n", res.JobId, res.Error)
			} else {
				log.Printf("Job %d success, returning : %v\n", res.JobId, res.Value)
			}
		}
		log.Println("Done fetching results for sending collaborator email notifications")
	}()

	return nil
}

func (uc *usecase) RemoveCollaborator(ctx context.Context, playlistId, userId int, req []collaborationModel.RemoveCollaboratorRequest) error {
	p, err := uc.playlistRepo.GetPlaylistById(ctx, userId)
	if err != nil {
		return err
	}

	// only the owner of the playlist can remove collaborator
	if p.OwnerId != userId {
		return collabErr.ErrRemoveCollaboratorNotPlaylistOwner
	}

	err = uc.collabRepo.RemoveCollaborator(ctx, playlistId, req)

	return err
}
