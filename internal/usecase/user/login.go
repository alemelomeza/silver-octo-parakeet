package userusecase

import (
	"context"
	"errors"

	"alemelomeza/silver-octo-parakeet/internal/repository"
	"alemelomeza/silver-octo-parakeet/internal/service/auth"
)

var ErrInvalidCredentials = errors.New("invalid username or password")
var ErrMustChangePassword = errors.New("user must change password")

type LoginUseCase struct {
	UserRepo repository.UserRepository
	Auth     auth.Service
}

func NewLoginUseCase(repo repository.UserRepository, authSvc auth.Service) *LoginUseCase {
	return &LoginUseCase{UserRepo: repo, Auth: authSvc}
}

func (uc *LoginUseCase) Execute(ctx context.Context, username, password string) (string, error) {
	u, err := uc.UserRepo.FindByUsername(ctx, username)
	if err != nil {
		return "", ErrInvalidCredentials
	}

	if !uc.Auth.CheckPassword(password, u.PasswordHash) {
		return "", ErrInvalidCredentials
	}

	if u.MustChangePwd {
		token, _ := uc.Auth.GenerateToken(u.ID, string(u.Role))
		return token, ErrMustChangePassword
	}

	token, err := uc.Auth.GenerateToken(u.ID, string(u.Role))
	return token, err
}
