package user

import (
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgconn"
	"github.com/medant81/url_library/internal/storage/postgre"
	"github.com/medant81/url_library/utils"
	"log/slog"
)

type repository struct {
	client postgre.Client
	log    *slog.Logger
}

func (r repository) Create(ctx context.Context, u *User) error {

	q := `
			insert into 
			    users(user_name, first_name, last_name, email, password)
			values
			    ($1, $2, $3, $4, $5)`

	r.log.InfoContext(ctx, fmt.Sprintf("SQL Query: %s", utils.FormatQuery(q)))

	if err := r.client.QueryRow(ctx, q, u.UserName, u.FirstName, u.LastName, u.Email, u.Password).Scan(&u.Id); err != nil {
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

func (r repository) FindOne(ctx context.Context, userName string) (User, error) {

	q := `
			select
				id,
				user_name,
				first_name,
				last_name,
				email,
				password
			from public.users
			where
				user_name = $1`

	r.log.InfoContext(ctx, fmt.Sprintf("SQL Query: %s", utils.FormatQuery(q)))

	var u User
	if err := r.client.QueryRow(ctx, q, userName).Scan(&u.Id, &u.UserName, &u.FirstName, &u.LastName, &u.Email, &u.Password); err != nil {
		return User{}, err
	}

	return u, nil
}

func (r repository) Update(ctx context.Context, u User) error {

	q := `
			update public.users
			set
				user_name = $2,
				first_name = $3,
				last_name = $4,
				email = $5,
				password = $6
			where id = $1`

	r.log.InfoContext(ctx, fmt.Sprintf("SQL Query: %s", utils.FormatQuery(q)))

	var count int
	if err := r.client.QueryRow(ctx, q, u.Id, u.UserName, u.FirstName, u.LastName, u.Email, u.Password).Scan(&count); err != nil {
		return err
	}

	return nil
}

func (r repository) Delete(ctx context.Context, userName string) error {

	q := `
			delete from public.users
			where user_name = $1
			returning id`

	r.log.InfoContext(ctx, fmt.Sprintf("SQL Query: %s", utils.FormatQuery(q)))

	var idDel int
	if err := r.client.QueryRow(ctx, q, userName).Scan(&idDel); err != nil {
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
