package playlist

import "errors"

var (
	ErrNotPlaylistOwner = errors.New("can't access the playlist because you're not the owner of the playlist")
	ErrPlaylistNotFound = errors.New("playlist not found")
)
