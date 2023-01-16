package models

import "time"

type Movie struct {
	ID           int       `json:"ID"`
	Title        string    `json:"title"`
	ReleaseDate  time.Time `json:"release_date"`
	Runtime      int       `json:"runtime"`
	MPAARRating  string    `json:"MPAAR_rating"`
	Description  string    `json:"description"`
	Image        string    `json:"image"`
	CreatedAt    time.Time `json:"-"`
	UpdatedField time.Time `json:"-"`
}
