package memory

import (
	"context"
	"errors"
	"sync"

	"alemelomeza/silver-octo-parakeet/internal/domain/user"

	"github.com/google/uuid"
)

var ErrUserNotFound = errors.New("user not found")

type UserRepositoryMemory struct {
	mu    sync.RWMutex
	items map[string]*user.User
}

func NewUserRepositoryMemory() *UserRepositoryMemory {
	return &UserRepositoryMemory{
		items: make(map[string]*user.User),
	}
}

func (r *UserRepositoryMemory) Create(ctx context.Context, u *user.User) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if u.ID == "" {
		u.ID = uuid.NewString()
	}

	copy := *u
	r.items[u.ID] = &copy

	return nil
}

func (r *UserRepositoryMemory) Update(ctx context.Context, u *user.User) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.items[u.ID]; !exists {
		return ErrUserNotFound
	}

	copy := *u
	r.items[u.ID] = &copy

	return nil
}

func (r *UserRepositoryMemory) Delete(ctx context.Context, id string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, ok := r.items[id]; !ok {
		return ErrUserNotFound
	}

	delete(r.items, id)
	return nil
}

func (r *UserRepositoryMemory) FindByID(ctx context.Context, id string) (*user.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	u, ok := r.items[id]
	if !ok {
		return nil, ErrUserNotFound
	}

	copy := *u
	return &copy, nil
}

func (r *UserRepositoryMemory) FindByUsername(ctx context.Context, username string) (*user.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	for _, u := range r.items {
		if u.Username == username {
			copy := *u
			return &copy, nil
		}
	}

	return nil, ErrUserNotFound
}

func (r *UserRepositoryMemory) List(ctx context.Context) ([]*user.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	list := make([]*user.User, 0, len(r.items))
	for _, u := range r.items {
		copy := *u
		list = append(list, &copy)
	}

	return list, nil
}
