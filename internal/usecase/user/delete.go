package userusecase

import (
    "context"

    "alemelomeza/silver-octo-parakeet/internal/domain/user"
    "alemelomeza/silver-octo-parakeet/internal/repository"
)

type DeleteUserUseCase struct {
    UserRepo repository.UserRepository
}

func NewDeleteUserUseCase(repo repository.UserRepository) *DeleteUserUseCase {
    return &DeleteUserUseCase{UserRepo: repo}
}

func (uc *DeleteUserUseCase) Execute(ctx context.Context, requesterRole user.Role, id string) error {
    if requesterRole != user.RoleAdmin {
        return user.ErrUnauthorized
    }
    return uc.UserRepo.Delete(ctx, id)
}
