package book

type CreateBookDTO struct {
	Name     string `json:"name"`
	Year     int    `json:"year"`
	AuthorID int    `json:"author_id"`
}
