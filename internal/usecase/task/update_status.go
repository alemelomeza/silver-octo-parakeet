package taskusecase

import (
    "context"
    "errors"
    "time"

    "alemelomeza/silver-octo-parakeet/internal/domain/task"
    "alemelomeza/silver-octo-parakeet/internal/domain/user"
    "alemelomeza/silver-octo-parakeet/internal/repository"
)

var ErrTaskExpired = errors.New("cannot update status of an expired task")

type UpdateTaskStatusUseCase struct {
    TaskRepo repository.TaskRepository
}

func NewUpdateTaskStatusUseCase(repo repository.TaskRepository) *UpdateTaskStatusUseCase {
    return &UpdateTaskStatusUseCase{TaskRepo: repo}
}

func (uc *UpdateTaskStatusUseCase) Execute(ctx context.Context, role user.Role, userID string, taskID string, newStatus task.Status) error {

    if role != user.RoleExecutor {
        return user.ErrUnauthorized
    }

    t, err := uc.TaskRepo.FindByID(ctx, taskID)
    if err != nil {
        return err
    }

    if t.AssignedTo != userID {
        return user.ErrUnauthorized
    }

    if time.Now().After(t.DueDate) {
        return ErrTaskExpired
    }

    t.Status = newStatus

    return uc.TaskRepo.Update(ctx, t)
}
