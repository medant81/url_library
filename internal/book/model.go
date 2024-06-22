package book

import "github.com/medant81/url_library/internal/author"

type Book struct {
	Id     int           `json:"id"`
	Name   string        `json:"name"`
	Year   int           `json:"year"`
	Isbn   string        `json:"isbn"`
	Author author.Author `json:"author"`
}
