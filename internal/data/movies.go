package data

import (
	"time"
)

type Movie struct {
	ID        int       `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	Title     string    `json:"title"`
	Year      int       `json:"year"`
	Runtime   int       `json:"runtime"`
	Genres    []string  `json:"genres"`
	Version   int       `json:"version"`
}
