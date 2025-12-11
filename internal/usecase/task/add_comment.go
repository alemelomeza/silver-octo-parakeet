package taskusecase

import (
	"context"
	"errors"
	"time"

	"alemelomeza/silver-octo-parakeet/internal/domain/user"
	"alemelomeza/silver-octo-parakeet/internal/repository"
)

var ErrTaskNotExpired = errors.New("can only comment on expired tasks")

type AddCommentUseCase struct {
	TaskRepo repository.TaskRepository
}

func NewAddCommentUseCase(repo repository.TaskRepository) *AddCommentUseCase {
	return &AddCommentUseCase{TaskRepo: repo}
}

func (uc *AddCommentUseCase) Execute(ctx context.Context, role user.Role, userID, taskID, comment string) error {

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

	if time.Now().Before(t.DueDate) {
		return ErrTaskNotExpired
	}

	t.Comments = append(t.Comments, comment)

	return uc.TaskRepo.Update(ctx, t)
}
