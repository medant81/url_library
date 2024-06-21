package handlers

import (
	"github.com/go-chi/chi/v5"
	"net/http"
	"strconv"
	"url_library/internal/author"
	"url_library/internal/book"
	"url_library/version"
)

type Application struct {
	RBook   book.Repository
	RAuthor author.Repository
}

func (app *Application) Home(w http.ResponseWriter, r *http.Request) {

	var homeJSON = struct {
		Status  string `json:"status"`
		Message string `json:"message"`
		Version string `json:"version"`
	}{
		Status:  "active",
		Message: "App start",
		Version: version.Version(),
	}

	_ = app.writeJSON(w, http.StatusOK, homeJSON)
}

func (app *Application) AddBook(w http.ResponseWriter, r *http.Request) {

	var b book.Book

	err := app.readJSON(w, r, &b)
	if err != nil {
		_ = app.errorJSON(w, err, http.StatusInternalServerError)
		return
	}

	err = app.RBook.Create(r.Context(), &b)
	if err != nil {
		_ = app.errorJSON(w, err, http.StatusInternalServerError)
		return
	}

	_ = app.writeJSON(w, http.StatusOK, b)
}

func (app *Application) AllBooks(w http.ResponseWriter, r *http.Request) {

	bs, err := app.RBook.FindAll(r.Context())
	if err != nil {
		_ = app.errorJSON(w, err, http.StatusInternalServerError)
		return
	}

	_ = app.writeJSON(w, http.StatusOK, bs)
}

func (app *Application) OneBook(w http.ResponseWriter, r *http.Request) {

	idString := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		_ = app.errorJSON(w, err, http.StatusBadRequest)
	}

	b, err := app.RBook.FindOne(r.Context(), id)
	if err != nil {
		_ = app.errorJSON(w, err, http.StatusInternalServerError)
		return
	}

	_ = app.writeJSON(w, http.StatusOK, b)
}

func (app *Application) UpdateBook(w http.ResponseWriter, r *http.Request) {

	idString := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		_ = app.errorJSON(w, err, http.StatusBadRequest)
	}

	var b book.Book
	err = app.readJSON(w, r, &b)
	if err != nil {
		_ = app.errorJSON(w, err, http.StatusBadRequest)
	}
	b.Id = id

	err = app.RBook.Update(r.Context(), &b)
	if err != nil {
		_ = app.errorJSON(w, err, http.StatusInternalServerError)
		return
	}

	_ = app.writeJSON(w, http.StatusOK, b)
}

func (app *Application) DelBook(w http.ResponseWriter, r *http.Request) {

	idString := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		_ = app.errorJSON(w, err, http.StatusBadRequest)
	}

	err = app.RBook.Delete(r.Context(), id)
	if err != nil {
		_ = app.errorJSON(w, err, http.StatusBadRequest)
	}

	_ = app.writeJSON(w, http.StatusOK, struct {
		Message string `json:"message"`
	}{
		Message: "",
	})
}

func (app *Application) UpdateBookWithAuthor(w http.ResponseWriter, r *http.Request) {

	idBookString := chi.URLParam(r, "book_id")
	idBook, err := strconv.Atoi(idBookString)
	if err != nil {
		_ = app.errorJSON(w, err, http.StatusBadRequest)
	}

	idAuthorString := chi.URLParam(r, "author_id")
	idAuthor, err := strconv.Atoi(idAuthorString)
	if err != nil {
		_ = app.errorJSON(w, err, http.StatusBadRequest)
	}

	var b book.Book
	err = app.readJSON(w, r, &b)
	if err != nil {
		_ = app.errorJSON(w, err, http.StatusBadRequest)
	}
	b.Id = idBook
	b.Author.Id = idAuthor

	err = app.RBook.UpdateWithAuthor(r.Context(), &b)
	if err != nil {
		_ = app.errorJSON(w, err, http.StatusInternalServerError)
		return
	}

	_ = app.writeJSON(w, http.StatusOK, b)
}

func (app *Application) AddAuthor(w http.ResponseWriter, r *http.Request) {

	var a author.Author

	err := app.readJSON(w, r, &a)
	if err != nil {
		_ = app.errorJSON(w, err, http.StatusInternalServerError)
		return
	}

	err = app.RAuthor.Create(r.Context(), &a)
	if err != nil {
		_ = app.errorJSON(w, err, http.StatusInternalServerError)
		return
	}

	_ = app.writeJSON(w, http.StatusOK, a)
}

func (app *Application) AllAuthors(w http.ResponseWriter, r *http.Request) {

	as, err := app.RAuthor.FindAll(r.Context())
	if err != nil {
		_ = app.errorJSON(w, err, http.StatusInternalServerError)
		return
	}

	_ = app.writeJSON(w, http.StatusOK, as)
}

func (app *Application) OneAuthor(w http.ResponseWriter, r *http.Request) {

	idString := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		_ = app.errorJSON(w, err, http.StatusBadRequest)
	}

	a, err := app.RAuthor.FindOne(r.Context(), id)
	if err != nil {
		_ = app.errorJSON(w, err, http.StatusInternalServerError)
		return
	}

	_ = app.writeJSON(w, http.StatusOK, a)
}

func (app *Application) UpdateAuthor(w http.ResponseWriter, r *http.Request) {

	idString := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		_ = app.errorJSON(w, err, http.StatusBadRequest)
	}

	var a author.Author
	err = app.readJSON(w, r, &a)
	if err != nil {
		_ = app.errorJSON(w, err, http.StatusBadRequest)
	}
	a.Id = id

	err = app.RAuthor.Update(r.Context(), &a)
	if err != nil {
		_ = app.errorJSON(w, err, http.StatusInternalServerError)
		return
	}

	_ = app.writeJSON(w, http.StatusOK, a)
}

func (app *Application) DelAuthor(w http.ResponseWriter, r *http.Request) {

	idString := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		_ = app.errorJSON(w, err, http.StatusBadRequest)
	}

	err = app.RAuthor.Delete(r.Context(), id)
	if err != nil {
		_ = app.errorJSON(w, err, http.StatusBadRequest)
	}

	_ = app.writeJSON(w, http.StatusOK, struct {
		Message string `json:"message"`
	}{
		Message: "",
	})
}
