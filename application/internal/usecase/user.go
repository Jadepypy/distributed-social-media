package usecase

import (
	"context"
	"github.com/Jadepypy/distributed-social-media/application/internal/domain"
	"github.com/Jadepypy/distributed-social-media/application/internal/repository"
	"golang.org/x/crypto/bcrypt"
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

func hashPassword(password string) ([]byte, error) {
	pwd := []byte(password)
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.DefaultCost)
	if err != nil {
		return []byte{}, err
	}
	return hash, nil
}

func (u *userUseCase) SignUp(ctx context.Context, user domain.User) error {
	hashedPassword, err := hashPassword(user.Password)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)

	err = u.repo.CreateUser(ctx, user)
	if err != nil {
		return err
	}

	return nil
}

func (u *userUseCase) LogIn(ctx context.Context, user domain.User) error {
	existingUser, err := u.repo.GetUserByEmail(ctx, user.Email)
	if err != nil {
		return err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(existingUser.Password), []byte(user.Password)); err != nil {
		return err
	}

	return nil
}

func (u *userUseCase) Edit(ctx context.Context, user domain.User) error {
	email := ctx.Value("user_id").(string)
	user.Email = email
	return u.repo.UpdateUser(ctx, user)
}

func (u *userUseCase) GetProfile(ctx context.Context, user domain.User) (domain.User, error) {
	//TODO implement me
	panic("implement me")
}
