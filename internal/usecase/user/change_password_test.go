package userusecase_test

import (
    "context"
    "testing"

    mem "alemelomeza/silver-octo-parakeet/internal/infrastructure/memory"
    "alemelomeza/silver-octo-parakeet/internal/service/auth"
    userusecase "alemelomeza/silver-octo-parakeet/internal/usecase/user"
    "alemelomeza/silver-octo-parakeet/internal/domain/user"
)

func TestChangePassword(t *testing.T) {
    repo := mem.NewUserRepositoryMemory()
    authSvc := auth.NewJWTService("secret", 24)

    u := &user.User{
        ID:           "10",
        Username:     "carlos",
        PasswordHash: authSvc.HashPassword("oldpass"),
    }
    repo.Create(context.Background(), u)

    uc := userusecase.NewChangePasswordUseCase(repo, authSvc)

    err := uc.Execute(context.Background(), "10", "oldpass", "newpass")
    if err != nil {
        t.Fatalf("unexpected error: %v", err)
    }

    updated, _ := repo.FindByID(context.Background(), "10")
    if !authSvc.CheckPassword("newpass", updated.PasswordHash) {
        t.Fatalf("password not updated properly")
    }
}
