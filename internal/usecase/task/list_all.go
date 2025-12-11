package taskusecase

import (
    "context"

    "alemelomeza/silver-octo-parakeet/internal/domain/user"
    "alemelomeza/silver-octo-parakeet/internal/repository"
    "alemelomeza/silver-octo-parakeet/internal/domain/task"
)

type ListAllTasksUseCase struct {
    TaskRepo repository.TaskRepository
}

func NewListAllTasksUseCase(repo repository.TaskRepository) *ListAllTasksUseCase {
    return &ListAllTasksUseCase{TaskRepo: repo}
}

func (uc *ListAllTasksUseCase) Execute(ctx context.Context, role user.Role) ([]*task.Task, error) {
    if role != user.RoleAuditor {
        return nil, user.ErrUnauthorized
    }
    return uc.TaskRepo.List(ctx)
}
