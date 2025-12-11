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

func TestUpdateTaskStatus(t *testing.T) {
    taskRepo := mem.NewTaskRepositoryMemory()

    t1 := &task.Task{
        ID:         "t1",
        AssignedTo: "exec1",
        Status:     task.StatusAssigned,
        DueDate:    time.Now().Add(1 * time.Hour),
    }

    taskRepo.Create(context.Background(), t1)

    uc := taskusecase.NewUpdateTaskStatusUseCase(taskRepo)

    err := uc.Execute(
        context.Background(),
        user.RoleExecutor,
        "exec1",
        "t1",
        task.StatusInProcess,
    )

    if err != nil {
        t.Fatalf("unexpected error: %v", err)
    }

    updated, _ := taskRepo.FindByID(context.Background(), "t1")
    if updated.Status != task.StatusInProcess {
        t.Fatalf("status was not updated")
    }
}
