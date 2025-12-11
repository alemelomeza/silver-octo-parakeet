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

func TestAddCommentToExpiredTask(t *testing.T) {
    taskRepo := mem.NewTaskRepositoryMemory()

    t1 := &task.Task{
        ID:         "t99",
        AssignedTo: "exec100",
        Status:     task.StatusInProcess,
        DueDate:    time.Now().Add(-1 * time.Hour), // expired
        Comments:   []string{},
    }

    taskRepo.Create(context.Background(), t1)

    uc := taskusecase.NewAddCommentUseCase(taskRepo)

    err := uc.Execute(
        context.Background(),
        user.RoleExecutor,
        "exec100",
        "t99",
        "Comment OK",
    )

    if err != nil {
        t.Fatalf("unexpected error: %v", err)
    }

    updated, _ := taskRepo.FindByID(context.Background(), "t99")
    if len(updated.Comments) != 1 {
        t.Fatalf("comment not added properly")
    }
}
