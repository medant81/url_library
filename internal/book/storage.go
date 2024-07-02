package book

import "context"

type Repository interface {
	Create(ctx context.Context, b *Book) error
	UpdateWithAuthor(ctx context.Context, b *Book) error
	FindAll(ctx context.Context) (b []Book, err error)
	FindOne(ctx context.Context, id int) (Book, error)
	Update(ctx context.Context, b *Book) error
	Delete(ctx context.Context, id int) error
}
