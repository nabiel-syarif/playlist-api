package collaboration

import "errors"

var (
	ErrAddCollaboratorNotPlaylistOwner = errors.New("only owner of the playlist can add as collaborator")
	ErrRemoveCollaboratorNotPlaylistOwner = errors.New("only owner of the playlist can remove collaborator")
)
