package taskusecase

import (
    "context"

    "alemelomeza/silver-octo-parakeet/internal/domain/user"
    "alemelomeza/silver-octo-parakeet/internal/repository"
    "alemelomeza/silver-octo-parakeet/internal/domain/task"
)

type ListMyTasksUseCase struct {
    TaskRepo repository.TaskRepository
}

func NewListMyTasksUseCase(repo repository.TaskRepository) *ListMyTasksUseCase {
    return &ListMyTasksUseCase{TaskRepo: repo}
}

func (uc *ListMyTasksUseCase) Execute(ctx context.Context, userID string, role user.Role) ([]*task.Task, error) {
    if role != user.RoleExecutor {
        return nil, user.ErrUnauthorized
    }

    return uc.TaskRepo.FindByUser(ctx, userID)
}
