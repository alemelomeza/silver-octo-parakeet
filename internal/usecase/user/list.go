package userusecase

import (
    "context"

    "alemelomeza/silver-octo-parakeet/internal/domain/user"
    "alemelomeza/silver-octo-parakeet/internal/repository"
)

type ListUsersUseCase struct {
    UserRepo repository.UserRepository
}

func NewListUsersUseCase(repo repository.UserRepository) *ListUsersUseCase {
    return &ListUsersUseCase{UserRepo: repo}
}

func (uc *ListUsersUseCase) Execute(ctx context.Context, requesterRole user.Role) ([]*user.User, error) {
    if requesterRole != user.RoleAdmin {
        return nil, user.ErrUnauthorized
    }
    return uc.UserRepo.List(ctx)
}
