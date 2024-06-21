package author

import (
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgconn"
	"log/slog"
	"url_library/internal/storage/postgre"
	"url_library/utils"
)

type repository struct {
	client postgre.Client
	log    *slog.Logger
}

func (r repository) Create(ctx context.Context, author *Author) error {

	q := `
		INSERT INTO authors 
		    (first_name, last_name, biography, birthday) 
		VALUES 
		       ($1, $2, $3, $4) 
		RETURNING id
	`
	r.log.InfoContext(ctx, fmt.Sprintf("SQL Query: %s", utils.FormatQuery(q)))
	if err := r.client.QueryRow(ctx, q, author.FirstName, author.LastName, author.Biography, author.Birthday).Scan(&author.Id); err != nil {
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

func (r repository) FindAll(ctx context.Context) (u []Author, err error) {
	q := `select
				id,
				first_name,
				last_name,
				biography,
				birthday
			from
				public.authors`
	r.log.InfoContext(ctx, fmt.Sprintf("SQL Query: %s", utils.FormatQuery(q)))

	rows, err := r.client.Query(ctx, q)
	if err != nil {
		return nil, err
	}

	authors := make([]Author, 0)

	for rows.Next() {
		var athr Author
		err = rows.Scan(&athr.Id, &athr.FirstName, &athr.LastName, &athr.Biography, &athr.Birthday)
		if err != nil {
			return nil, err
		}
		authors = append(authors, athr)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return authors, nil
}

func (r repository) FindOne(ctx context.Context, id int) (Author, error) {

	q := `select
				id,
				first_name,
				last_name,
				biography,
				birthday
			from
				public.authors
			where
				id = $1`

	r.log.InfoContext(ctx, fmt.Sprintf("SQL Query: %s", utils.FormatQuery(q)))

	var athr Author
	if err := r.client.QueryRow(ctx, q, id).Scan(&athr.Id, &athr.FirstName, &athr.LastName, &athr.Biography, &athr.Birthday); err != nil {
		return Author{}, err
	}

	return athr, nil
}

func (r repository) Update(ctx context.Context, athr *Author) error {

	q := `update public.authors
			set
				first_name = $2,
				last_name = $3,
				biography = $4,
				birthday = $5
			where id = $1`

	r.log.InfoContext(ctx, fmt.Sprintf("SQL Query: %s", utils.FormatQuery(q)))

	var count int
	if err := r.client.QueryRow(ctx, q, athr.Id, athr.FirstName, athr.LastName, athr.Biography, athr.Birthday).Scan(&count); err != nil {
		return err
	}

	return nil
}

func (r repository) Delete(ctx context.Context, id int) error {

	q := `
			delete from public.authors
			where id = $1
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
