package models

type Movie struct {
	Id          int
	Title       string
	Rating      int
	Description string
	Director    string
	ReleaseYear int
	Genres      []Genre
	TrailerUrl  string
	PosterUrl   string
	IsWatched   bool
}
