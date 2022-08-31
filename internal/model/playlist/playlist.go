package playlist

import (
	"time"

	modelSong "github.com/nabiel-syarif/playlist-api/internal/model/song"
)

type Playlist struct {
	Id        int       `json:"id"`
	Name      string    `json:"name"`
	OwnerId   int       `json:"owner_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type PlaylistAggregated struct {
	Id        int              `json:"id"`
	Name      string           `json:"name"`
	Songs     []modelSong.Song `json:"songs"`
	CreatedAt time.Time        `json:"created_at"`
	UpdatedAt time.Time        `json:"updated_at"`
}

type AddPlaylistRequest struct {
	Name  string `json:"name"`
	Owner int    `json:"owner_id"`
}

type UpdatePlaylistRequest struct {
	Id    int    `json:"playlist_id"`
	Name  string `json:"name"`
	Owner int    `json:"owner_id"`
}
