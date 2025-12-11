package repository

import (
    "context"
    "alemelomeza/silver-octo-parakeet/internal/domain/task"
)

type TaskRepository interface {
    Create(ctx context.Context, t *task.Task) error
    Update(ctx context.Context, t *task.Task) error
    Delete(ctx context.Context, id string) error
    
    FindByID(ctx context.Context, id string) (*task.Task, error)
    FindByUser(ctx context.Context, userID string) ([]*task.Task, error)
    List(ctx context.Context) ([]*task.Task, error)
}
