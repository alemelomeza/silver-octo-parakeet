package taskusecase

import (
    "context"
    "errors"

    "alemelomeza/silver-octo-parakeet/internal/domain/task"
    "alemelomeza/silver-octo-parakeet/internal/domain/user"
    "alemelomeza/silver-octo-parakeet/internal/repository"
)

var ErrTaskNotDeletable = errors.New("only tasks in 'Asignado' status can be deleted")

type DeleteTaskUseCase struct {
    TaskRepo repository.TaskRepository
}

func NewDeleteTaskUseCase(repo repository.TaskRepository) *DeleteTaskUseCase {
    return &DeleteTaskUseCase{TaskRepo: repo}
}

func (uc *DeleteTaskUseCase) Execute(ctx context.Context, role user.Role, id string) error {
    if role != user.RoleAdmin {
        return user.ErrUnauthorized
    }

    t, err := uc.TaskRepo.FindByID(ctx, id)
    if err != nil {
        return err
    }

    if t.Status != task.StatusAssigned {
        return ErrTaskNotDeletable
    }

    return uc.TaskRepo.Delete(ctx, id)
}
