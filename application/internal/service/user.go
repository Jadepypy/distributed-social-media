package service

import (
	"context"
	"errors"
	"github.com/Jadepypy/distributed-social-media/application/internal/domain"
	"github.com/Jadepypy/distributed-social-media/application/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

var (
	ErrUserNotFound    = errors.New("user not found")
	ErrorWrongPassword = errors.New("wrong password")
)

type UserService interface {
	SignUp(ctx context.Context, user domain.User) error
	LogIn(ctx context.Context, user domain.User) error
	Edit(ctx context.Context, user domain.User) error
	GetProfile(ctx context.Context, email string) (domain.User, error)
}

type userService struct {
	repo repository.UserRepository
}

func NewUserUseCase(repo repository.UserRepository) UserService {
	return &userService{
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

func (u *userService) SignUp(ctx context.Context, user domain.User) error {
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

func (u *userService) LogIn(ctx context.Context, user domain.User) error {
	existingUser, err := u.repo.GetUserByEmail(ctx, user.Email)
	if err != nil {
		return ErrUserNotFound
	}

	if err = bcrypt.CompareHashAndPassword([]byte(existingUser.Password), []byte(user.Password)); err != nil {
		return ErrorWrongPassword
	}

	return nil
}

func (u *userService) Edit(ctx context.Context, user domain.User) error {
	return u.repo.UpdateUser(ctx, user)
}

func (u *userService) GetProfile(ctx context.Context, email string) (domain.User, error) {
	return u.repo.GetUserByEmail(ctx, email)
}
