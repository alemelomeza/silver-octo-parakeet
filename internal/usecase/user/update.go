package userusecase

import (
    "context"

    "alemelomeza/silver-octo-parakeet/internal/domain/user"
    "alemelomeza/silver-octo-parakeet/internal/repository"
)

type UpdateUserUseCase struct {
    UserRepo repository.UserRepository
}

func NewUpdateUserUseCase(repo repository.UserRepository) *UpdateUserUseCase {
    return &UpdateUserUseCase{UserRepo: repo}
}

func (uc *UpdateUserUseCase) Execute(ctx context.Context, requesterRole user.Role, u *user.User) error {
    if requesterRole != user.RoleAdmin {
        return user.ErrUnauthorized
    }
    return uc.UserRepo.Update(ctx, u)
}
