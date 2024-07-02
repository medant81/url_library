package handlers

import (
	"github.com/go-chi/chi/v5"

	"github.com/medant81/url_library/internal/author"
	"github.com/medant81/url_library/internal/book"
	"github.com/medant81/url_library/version"
	"net/http"
	"strconv"
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

// @Summary Create book
// @Tags books
// @Description create book
// @ID create-book
// @Accept  json
// @Produce  json
// @Param input body book.Book true "book info"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} JSONResponse
// @Failure 500 {object} JSONResponse
// @Failure default {object} JSONResponse
// @Router /books [post]
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

// @Summary Get all books
// @Tags books
// @Description get all books
// @ID get-books
// @Accept  json
// @Produce  json
// @Success 200 {object} []book.Book
// @Failure 400,404 {object} JSONResponse
// @Failure 500 {object} JSONResponse
// @Failure default {object} JSONResponse
// @Router /books [get]
func (app *Application) AllBooks(w http.ResponseWriter, r *http.Request) {

	bs, err := app.RBook.FindAll(r.Context())
	if err != nil {
		_ = app.errorJSON(w, err, http.StatusInternalServerError)
		return
	}

	_ = app.writeJSON(w, http.StatusOK, bs)
}

// @Summary Get book
// @Tags books
// @Description get book
// @ID get-book
// @Accept  json
// @Produce  json
// @Success 200 {object} book.Book
// @Failure 400,404 {object} JSONResponse
// @Failure 500 {object} JSONResponse
// @Failure default {object} JSONResponse
// @Router /books/:id [get]
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

// @Summary Update book
// @Tags books
// @Description update book
// @ID update-book
// @Accept  json
// @Produce  json
// @Success 200 {object} book.Book
// @Failure 400,404 {object} JSONResponse
// @Failure 500 {object} JSONResponse
// @Failure default {object} JSONResponse
// @Router /books/:id [put]
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

// @Summary Delete book
// @Tags books
// @Description delete book
// @ID delete-book
// @Accept  json
// @Produce  json
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} JSONResponse
// @Failure 500 {object} JSONResponse
// @Failure default {object} JSONResponse
// @Router /books/:id [delete]
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

// @Summary Update book with author
// @Tags books
// @Description update book with author
// @ID update-book-author
// @Accept  json
// @Produce  json
// @Success 200 {object} book.Book
// @Failure 400,404 {object} JSONResponse
// @Failure 500 {object} JSONResponse
// @Failure default {object} JSONResponse
// @Router /books/:id_book/authors/:id_athor [put]
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

// @Summary Create author
// @Tags authors
// @Description create author
// @ID create-author
// @Accept  json
// @Produce  json
// @Param input body author.Author true "author info"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} JSONResponse
// @Failure 500 {object} JSONResponse
// @Failure default {object} JSONResponse
// @Router /authors [post]
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

// @Summary Get all authors
// @Tags authors
// @Description get all authors
// @ID get-authors
// @Accept  json
// @Produce  json
// @Success 200 {object} []author.Author
// @Failure 400,404 {object} JSONResponse
// @Failure 500 {object} JSONResponse
// @Failure default {object} JSONResponse
// @Router /authors [get]
func (app *Application) AllAuthors(w http.ResponseWriter, r *http.Request) {

	as, err := app.RAuthor.FindAll(r.Context())
	if err != nil {
		_ = app.errorJSON(w, err, http.StatusInternalServerError)
		return
	}

	_ = app.writeJSON(w, http.StatusOK, as)
}

// @Summary Get author
// @Tags authors
// @Description get author
// @ID get-author
// @Accept  json
// @Produce  json
// @Success 200 {object} author.Author
// @Failure 400,404 {object} JSONResponse
// @Failure 500 {object} JSONResponse
// @Failure default {object} JSONResponse
// @Router /authors/:id [get]
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

// @Summary Update author
// @Tags authors
// @Description update author
// @ID update-author
// @Accept  json
// @Produce  json
// @Success 200 {object} author.Author
// @Failure 400,404 {object} JSONResponse
// @Failure 500 {object} JSONResponse
// @Failure default {object} JSONResponse
// @Router /authors/:id [put]
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

// @Summary Delete author
// @Tags authors
// @Description delete author
// @ID delete-author
// @Accept  json
// @Produce  json
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} JSONResponse
// @Failure 500 {object} JSONResponse
// @Failure default {object} JSONResponse
// @Router /authors/:id [delete]
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
