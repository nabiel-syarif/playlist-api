package song

import "time"

type Song struct {
	Id        int       `json:"id"`
	Title     string    `json:"title"`
	Performer string    `json:"performer"`
	Genre     string    `json:"genre"`
	Duration  float64   `json:"duration"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type InsertSongRequest struct {
	Title     string    `json:"title"`
	Performer string    `json:"performer"`
	Genre     string    `json:"genre"`
	Duration  float64   `json:"duration"`
}