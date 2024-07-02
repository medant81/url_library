package book

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/jackc/pgconn"
	"github.com/medant81/url_library/utils"
	"log/slog"

	"github.com/medant81/url_library/internal/storage/postgre"
)

type repository struct {
	client postgre.Client
	log    *slog.Logger
}

func (r repository) UpdateWithAuthor(ctx context.Context, b *Book) error {

	tx, err := r.client.Begin(ctx)

	if err != nil {
		return err
	}

	q := `update public.authors
			set
				first_name = $2,
				last_name = $3,
				biography = $4,
				birthday = $5
			where id = $1`

	r.log.InfoContext(ctx, fmt.Sprintf("SQL Query: %s", utils.FormatQuery(q)))

	var count int
	if err := r.client.QueryRow(ctx, q, b.Author.Id, b.Author.FirstName, b.Author.LastName, b.Author.Biography, b.Author.Birthday).Scan(&count); err != nil {
		_ = tx.Rollback(ctx)
		return err
	}

	q = `
			update public.books
			set
				name = $2,
				author_id = $3,
				year = $4,
				isbn = $5
			where id = $1`

	r.log.InfoContext(ctx, fmt.Sprintf("SQL Query: %s", utils.FormatQuery(q)))

	if err := r.client.QueryRow(ctx, q, b.Id, b.Name, b.Author.Id, b.Year, b.Isbn).Scan(&count); err != nil {
		_ = tx.Rollback(ctx)
		return err
	}

	if err = tx.Commit(ctx); err != nil {
		return err
	}

	return nil
}

func (r repository) Create(ctx context.Context, b *Book) error {

	q := `
			insert into books
    			(name, year, author_id, isbn)
			values
			    ($1, $2, $3, $4)
			returning id`

	r.log.InfoContext(ctx, fmt.Sprintf("SQL Query: %s", utils.FormatQuery(q)))

	if err := r.client.QueryRow(ctx, q, b.Name, b.Year, b.Author.Id, b.Isbn).Scan(&b.Id); err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			pgErr = err.(*pgconn.PgError)
			strErr := fmt.Sprintf("SQL Error: %s, Detail: %s, Where: %s, Code: %s, SQLState: %s", pgErr.Message, pgErr.Detail, pgErr.Where, pgErr.Code, pgErr.SQLState())
			newErr := fmt.Errorf(strErr)
			r.log.Error(strErr)
			return newErr
		}
		return err
	}
	return nil

}

func (r repository) FindAll(ctx context.Context) (b []Book, err error) {

	q := `
			select
				id,
				name,
				(select
					json_build_object(
						'id', authors.id
						,'first_name', authors.first_name
						,'last_name', authors.last_name
						,'birthday', to_char(authors.birthday, 'YYYY-MM-DD"T"HH24:MI:SS"Z"')
						) as author
					from public.authors as authors
					where authors.id = author_id) as author,
				year,
				isbn
			from
				public.books`

	r.log.InfoContext(ctx, fmt.Sprintf("SQL Query: %s", utils.FormatQuery(q)))

	rows, err := r.client.Query(ctx, q)
	if err != nil {
		return nil, err
	}

	books := make([]Book, 0)

	for rows.Next() {
		var bk Book
		var sAuthor string
		err = rows.Scan(&bk.Id, &bk.Name, &sAuthor, &bk.Year, &bk.Isbn)
		if err != nil {
			return nil, err
		}
		_ = json.Unmarshal([]byte(sAuthor), &bk.Author)
		books = append(books, bk)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return books, nil

}

func (r repository) FindOne(ctx context.Context, id int) (Book, error) {

	q := `
			select
				id,
				name,
				(select
					json_build_object(
						'id', authors.id
						,'first_name', authors.first_name
						,'last_name', authors.last_name
						,'birthday', to_char(authors.birthday, 'YYYY-MM-DD"T"HH24:MI:SS"Z"')
						) as author
					from public.authors as authors
					where authors.id = author_id) as author,
				year,
				isbn
			from
				public.books
			where
				id = $1`

	r.log.InfoContext(ctx, fmt.Sprintf("SQL Query: %s", utils.FormatQuery(q)))

	var bk Book
	var sAuthor string
	if err := r.client.QueryRow(ctx, q, id).Scan(&bk.Id, &bk.Name, &sAuthor, &bk.Year, &bk.Isbn); err != nil {
		return Book{}, err
	}

	_ = json.Unmarshal([]byte(sAuthor), &bk.Author)

	return bk, nil
}

func (r repository) Update(ctx context.Context, b *Book) error {

	q := `
			update public.books
			set
				name = $2,
				author_id = $3,
				year = $4,
				isbn = $5
			where id = $1`

	r.log.InfoContext(ctx, fmt.Sprintf("SQL Query: %s", utils.FormatQuery(q)))

	var count int
	if err := r.client.QueryRow(ctx, q, b.Id, b.Name, b.Author.Id, b.Year, b.Isbn).Scan(&count); err != nil {
		return err
	}

	return nil

}

func (r repository) Delete(ctx context.Context, id int) error {

	q := `
			delete from public.books
			where id =$1
			returning id`

	r.log.InfoContext(ctx, fmt.Sprintf("SQL Query: %s", utils.FormatQuery(q)))

	var idDel int
	if err := r.client.QueryRow(ctx, q, id).Scan(&idDel); err != nil {
		return err
	}

	return nil
}

func NewRepository(client postgre.Client, log *slog.Logger) Repository {
	return &repository{
		client: client,
		log:    log,
	}
}
