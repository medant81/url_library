package user

import "context"

type Repository interface {
	Create(ctx context.Context, user *User) error
	FindOne(ctx context.Context, userName string) (User, error)
	Update(ctx context.Context, user User) error
	Delete(ctx context.Context, userName string) error
}
