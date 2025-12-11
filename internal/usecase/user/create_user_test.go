package userusecase_test

import (
    "context"
    "testing"

    mem "alemelomeza/silver-octo-parakeet/internal/infrastructure/memory"
    "alemelomeza/silver-octo-parakeet/internal/service/auth"
    userusecase "alemelomeza/silver-octo-parakeet/internal/usecase/user"
    "alemelomeza/silver-octo-parakeet/internal/domain/user"
)

func TestCreateUserAsAdmin(t *testing.T) {
    repo := mem.NewUserRepositoryMemory()
    authSvc := auth.NewJWTService("secret", 24)

    uc := userusecase.NewCreateUserUseCase(repo, authSvc)

    created, err := uc.Execute(
        context.Background(),
        user.RoleAdmin,
        "newexec",
        user.RoleExecutor,
    )

    if err != nil {
        t.Fatalf("unexpected error: %v", err)
    }

    if created.Username != "newexec" {
        t.Fatalf("expected username newexec, got %v", created.Username)
    }
}

func TestCreateUserAsExecutorForbidden(t *testing.T) {
    repo := mem.NewUserRepositoryMemory()
    authSvc := auth.NewJWTService("secret", 24)

    uc := userusecase.NewCreateUserUseCase(repo, authSvc)

    _, err := uc.Execute(
        context.Background(),
        user.RoleExecutor,
        "test",
        user.RoleExecutor,
    )

    if err == nil {
        t.Fatalf("expected forbidden error, got nil")
    }
}
