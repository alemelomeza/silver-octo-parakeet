package userusecase

import (
    "context"
    "errors"

    "alemelomeza/silver-octo-parakeet/internal/domain/user"
    "alemelomeza/silver-octo-parakeet/internal/repository"
    "alemelomeza/silver-octo-parakeet/internal/service/auth"
)

var ErrForbiddenRole = errors.New("admin cannot create another admin user")

type CreateUserUseCase struct {
    UserRepo repository.UserRepository
    Auth     auth.Service
}

func NewCreateUserUseCase(repo repository.UserRepository, authSvc auth.Service) *CreateUserUseCase {
    return &CreateUserUseCase{UserRepo: repo, Auth: authSvc}
}

func (uc *CreateUserUseCase) Execute(ctx context.Context, creatorRole user.Role, username string, role user.Role) (*user.User, error) {

    if creatorRole != user.RoleAdmin {
        return nil, user.ErrUnauthorized
    }

    if role == user.RoleAdmin {
        return nil, ErrForbiddenRole
    }

    tempPassword := uc.Auth.GenerateTempPassword()
    hash := uc.Auth.HashPassword(tempPassword)

    newUser := &user.User{
        Username:      username,
        Role:          role,
        PasswordHash:  hash,
        MustChangePwd: true,
    }

    err := uc.UserRepo.Create(ctx, newUser)
    if err != nil {
        return nil, err
    }

    return newUser, nil
}
