package memory

import (
	"alemelomeza/silver-octo-parakeet/internal/domain/task"
	"context"
	"errors"
	"sync"

	"github.com/google/uuid"
)

var ErrTaskNotFound = errors.New("task not found")

type TaskRepositoryMemory struct {
	mu    sync.RWMutex
	items map[string]*task.Task
}

func NewTaskRepositoryMemory() *TaskRepositoryMemory {
	return &TaskRepositoryMemory{
		items: make(map[string]*task.Task),
	}
}

func (r *TaskRepositoryMemory) Create(ctx context.Context, t *task.Task) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if t.ID == "" {
		t.ID = uuid.NewString()
	}

	copy := *t
	r.items[t.ID] = &copy

	return nil
}

func (r *TaskRepositoryMemory) Update(ctx context.Context, t *task.Task) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	_, exists := r.items[t.ID]
	if !exists {
		return ErrTaskNotFound
	}

	copy := *t
	r.items[t.ID] = &copy

	return nil
}

func (r *TaskRepositoryMemory) Delete(ctx context.Context, id string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, ok := r.items[id]; !ok {
		return ErrTaskNotFound
	}

	delete(r.items, id)
	return nil
}

func (r *TaskRepositoryMemory) FindByID(ctx context.Context, id string) (*task.Task, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	t, ok := r.items[id]
	if !ok {
		return nil, ErrTaskNotFound
	}

	copy := *t
	return &copy, nil
}

func (r *TaskRepositoryMemory) FindByUser(ctx context.Context, userID string) ([]*task.Task, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	list := []*task.Task{}
	for _, t := range r.items {
		if t.AssignedTo == userID {
			copy := *t
			list = append(list, &copy)
		}
	}

	return list, nil
}

func (r *TaskRepositoryMemory) List(ctx context.Context) ([]*task.Task, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	list := make([]*task.Task, 0, len(r.items))
	for _, t := range r.items {
		copy := *t
		list = append(list, &copy)
	}

	return list, nil
}
