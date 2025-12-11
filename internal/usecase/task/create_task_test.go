package taskusecase_test

import (
    "context"
    "testing"
    "time"

    mem "alemelomeza/silver-octo-parakeet/internal/infrastructure/memory"
    taskusecase "alemelomeza/silver-octo-parakeet/internal/usecase/task"
    "alemelomeza/silver-octo-parakeet/internal/domain/task"
    "alemelomeza/silver-octo-parakeet/internal/domain/user"
)

func TestCreateTaskAsAdmin(t *testing.T) {
    taskRepo := mem.NewTaskRepositoryMemory()
    userRepo := mem.NewUserRepositoryMemory()

    // Add executor
    userRepo.Create(context.Background(), &user.User{
        ID:       "exec1",
        Role:     user.RoleExecutor,
        Username: "john",
    })

    uc := taskusecase.NewCreateTaskUseCase(taskRepo, userRepo)

    due := time.Now().Add(24 * time.Hour)

    newTask, err := uc.Execute(
        context.Background(),
        user.RoleAdmin,
        "Task A",
        "Description",
        due,
        "exec1",
    )

    if err != nil {
        t.Fatalf("unexpected error: %v", err)
    }

    if newTask.Status != task.StatusAssigned {
        t.Fatalf("expected status Assigned, got %v", newTask.Status)
    }
}
