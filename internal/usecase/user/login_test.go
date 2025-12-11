package userusecase_test

import (
    "context"
    "testing"

    mem "alemelomeza/silver-octo-parakeet/internal/infrastructure/memory"
    "alemelomeza/silver-octo-parakeet/internal/service/auth"
    userusecase "alemelomeza/silver-octo-parakeet/internal/usecase/user"
    "alemelomeza/silver-octo-parakeet/internal/domain/user"
)

func TestLoginSuccess(t *testing.T) {
    repo := mem.NewUserRepositoryMemory()
    authSvc := auth.NewJWTService("secret", 24)
    uc := userusecase.NewLoginUseCase(repo, authSvc)

    u := &user.User{
        ID:            "1",
        Username:      "john",
        Role:          user.RoleExecutor,
        PasswordHash:  authSvc.HashPassword("1234"),
        MustChangePwd: false,
    }
    repo.Create(context.Background(), u)

    token, err := uc.Execute(context.Background(), "john", "1234")
    if err != nil {
        t.Fatalf("expected no error, got %v", err)
    }

    if token == "" {
        t.Fatalf("expected token, got empty string")
    }
}

func TestLoginInvalidPassword(t *testing.T) {
    repo := mem.NewUserRepositoryMemory()
    authSvc := auth.NewJWTService("secret", 24)
    uc := userusecase.NewLoginUseCase(repo, authSvc)

    u := &user.User{
        ID:            "1",
        Username:      "jane",
        Role:          user.RoleExecutor,
        PasswordHash:  authSvc.HashPassword("abcd"),
        MustChangePwd: false,
    }
    repo.Create(context.Background(), u)

    _, err := uc.Execute(context.Background(), "jane", "wrong")
    if err == nil {
        t.Fatalf("expected error, got nil")
    }
}
