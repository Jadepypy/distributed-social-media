package usecase

import (
	"context"
	"github.com/Jadepypy/distributed-social-media/application/internal/domain"
	"github.com/Jadepypy/distributed-social-media/application/internal/repository"
)

type UserUseCase interface {
	SignUp(ctx context.Context, user domain.User) error
	LogIn(ctx context.Context, user domain.User) error
	Edit(ctx context.Context, user domain.User) error
	GetProfile(ctx context.Context, user domain.User) (domain.User, error)
}

type userUseCase struct {
	repo repository.UserRepository
}

func NewUserUseCase(repo repository.UserRepository) UserUseCase {
	return &userUseCase{
		repo: repo,
	}
}
func (u *userUseCase) SignUp(ctx context.Context, user domain.User) error {
	//TODO implement me
	panic("implement me")
}

func (u *userUseCase) LogIn(ctx context.Context, user domain.User) error {
	//TODO implement me
	panic("implement me")
}

func (u *userUseCase) Edit(ctx context.Context, user domain.User) error {
	//TODO implement me
	panic("implement me")
}

func (u *userUseCase) GetProfile(ctx context.Context, user domain.User) (domain.User, error) {
	//TODO implement me
	panic("implement me")
}
