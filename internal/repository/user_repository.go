package repository

import (
	"alemelomeza/silver-octo-parakeet/internal/domain/user"
	"context"
)

type UserRepository interface {
    Create(ctx context.Context, u *user.User) error
    Update(ctx context.Context, u *user.User) error
    Delete(ctx context.Context, id string) error

    FindByID(ctx context.Context, id string) (*user.User, error)
    FindByUsername(ctx context.Context, username string) (*user.User, error)
    List(ctx context.Context) ([]*user.User, error)
}
