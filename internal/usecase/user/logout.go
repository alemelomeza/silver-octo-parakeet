package userusecase

import "context"

type LogoutUseCase struct{}

func NewLogoutUseCase() *LogoutUseCase {
	return &LogoutUseCase{}
}

func (uc *LogoutUseCase) Execute(ctx context.Context, token string) error {
	return nil
}
