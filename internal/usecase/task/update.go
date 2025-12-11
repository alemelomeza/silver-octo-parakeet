package taskusecase

import (
    "context"
    "errors"

    "alemelomeza/silver-octo-parakeet/internal/domain/task"
    "alemelomeza/silver-octo-parakeet/internal/domain/user"
    "alemelomeza/silver-octo-parakeet/internal/repository"
)

var ErrTaskNotUpdatable = errors.New("only tasks in 'Asignado' status can be updated")

type UpdateTaskUseCase struct {
    TaskRepo repository.TaskRepository
}

func NewUpdateTaskUseCase(taskRepo repository.TaskRepository) *UpdateTaskUseCase {
    return &UpdateTaskUseCase{TaskRepo: taskRepo}
}

func (uc *UpdateTaskUseCase) Execute(ctx context.Context, role user.Role, t *task.Task) error {
    if role != user.RoleAdmin {
        return user.ErrUnauthorized
    }

    stored, err := uc.TaskRepo.FindByID(ctx, t.ID)
    if err != nil {
        return err
    }

    if stored.Status != task.StatusAssigned {
        return ErrTaskNotUpdatable
    }

    return uc.TaskRepo.Update(ctx, t)
}
