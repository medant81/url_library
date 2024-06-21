package book

import "url_library/internal/author"

type Book struct {
	Id     int           `json:"id"`
	Name   string        `json:"name"`
	Year   int           `json:"year"`
	Isbn   string        `json:"isbn"`
	Author author.Author `json:"author"`
}
