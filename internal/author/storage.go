package author

import "context"

type Repository interface {
	Create(ctx context.Context, a *Author) error
	FindAll(ctx context.Context) (u []Author, err error)
	FindOne(ctx context.Context, id int) (Author, error)
	Update(ctx context.Context, a *Author) error
	Delete(ctx context.Context, id int) error
}
