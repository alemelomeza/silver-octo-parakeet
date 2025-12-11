package taskusecase

import (
    "context"
    "errors"
    "time"

    "alemelomeza/silver-octo-parakeet/internal/domain/task"
    "alemelomeza/silver-octo-parakeet/internal/domain/user"
    "alemelomeza/silver-octo-parakeet/internal/repository"
)

var ErrInvalidAssignee = errors.New("task can only be assigned to an executor user")

type CreateTaskUseCase struct {
    TaskRepo repository.TaskRepository
    UserRepo repository.UserRepository
}

func NewCreateTaskUseCase(taskRepo repository.TaskRepository, userRepo repository.UserRepository) *CreateTaskUseCase {
    return &CreateTaskUseCase{
        TaskRepo: taskRepo,
        UserRepo: userRepo,
    }
}

func (uc *CreateTaskUseCase) Execute(
    ctx context.Context,
    creatorRole user.Role,
    title string,
    description string,
    dueDate time.Time,
    assignedTo string,
) (*task.Task, error) {

    if creatorRole != user.RoleAdmin {
        return nil, user.ErrUnauthorized
    }

    assignee, err := uc.UserRepo.FindByID(ctx, assignedTo)
    if err != nil || assignee.Role != user.RoleExecutor {
        return nil, ErrInvalidAssignee
    }

    newTask := &task.Task{
        Title:       title,
        Description: description,
        DueDate:     dueDate,
        AssignedTo:  assignedTo,
        Status:      task.StatusAssigned,
        Comments:    []string{},
    }

    if err := uc.TaskRepo.Create(ctx, newTask); err != nil {
        return nil, err
    }

    return newTask, nil
}
