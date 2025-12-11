package userusecase

import (
	"context"
	"errors"

	"alemelomeza/silver-octo-parakeet/internal/repository"
	"alemelomeza/silver-octo-parakeet/internal/service/auth"
)

var ErrOldPasswordIncorrect = errors.New("old password incorrect")

type ChangePasswordUseCase struct {
	UserRepo repository.UserRepository
	Auth     auth.Service
}

func NewChangePasswordUseCase(repo repository.UserRepository, authSvc auth.Service) *ChangePasswordUseCase {
	return &ChangePasswordUseCase{UserRepo: repo, Auth: authSvc}
}

func (uc *ChangePasswordUseCase) Execute(ctx context.Context, userID, oldPassword, newPassword string) error {

	u, err := uc.UserRepo.FindByID(ctx, userID)
	if err != nil {
		return err
	}

	if !uc.Auth.CheckPassword(oldPassword, u.PasswordHash) {
		return ErrOldPasswordIncorrect
	}

	u.PasswordHash = uc.Auth.HashPassword(newPassword)
	u.MustChangePwd = false

	return uc.UserRepo.Update(ctx, u)
}
