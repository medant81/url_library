package handlers

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
)

func Routers(app *Application) http.Handler {
	mux := chi.NewRouter()
	mux.Use(middleware.Recoverer)
	mux.Get("/", app.Home)
	mux.Post("/books", app.AddBook)
	mux.Get("/books", app.AllBooks)
	mux.Get("/books/{id}", app.OneBook)
	mux.Put("/books/{id}", app.UpdateBook)
	mux.Delete("/books/{id}", app.DelBook)
	mux.Post("/authors", app.AddAuthor)
	mux.Get("/authors", app.AllAuthors)
	mux.Get("/authors/{id}", app.OneAuthor)
	mux.Put("/authors/{id}", app.UpdateAuthor)
	mux.Delete("/authors/{id}", app.DelAuthor)
	mux.Put("/books/{book_id}/authors/{author_id}", app.UpdateBookWithAuthor)
	return mux
}
