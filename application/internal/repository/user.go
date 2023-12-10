package repository

import (
	"context"
	"github.com/Jadepypy/distributed-social-media/application/internal/domain"
	"github.com/Jadepypy/distributed-social-media/application/internal/repository/dao"
	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user domain.User) error
	GetUserByEmail(ctx context.Context, email string) (domain.User, error)
	UpdateUser(ctx context.Context, user domain.User) error
}

type userRepository struct {
	dao dao.UserDAO
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		dao: dao.NewUserDao(db),
	}
}

func (u *userRepository) CreateUser(ctx context.Context, user domain.User) error {
	return u.dao.Insert(ctx, dao.User{
		Email:    user.Email,
		Password: user.Password,
	})
}

func (u *userRepository) GetUserByEmail(ctx context.Context, email string) (domain.User, error) {
	user, err := u.dao.GetByEmail(ctx, email)
	if err != nil {
		return domain.User{}, err
	}
	return daoToDomain(user), nil
}

func (u *userRepository) UpdateUser(ctx context.Context, user domain.User) error {
	return u.dao.Update(ctx, user.Email, domainToDao(user))
}

func domainToDao(user domain.User) dao.User {
	u := dao.User{}
	if user.Birthday != "" {
		u.Birthday = &user.Birthday
	}
	if user.Nickname != "" {
		u.Nickname = &user.Nickname
	}
	if user.Intro != "" {
		u.Intro = &user.Intro
	}
	return u
}
func daoToDomain(user dao.User) domain.User {
	u := domain.User{
		Email:    user.Email,
		Password: user.Password,
	}
	if user.Birthday != nil {
		u.Birthday = *user.Birthday
	}
	if user.Nickname != nil {
		u.Nickname = *user.Nickname
	}
	if user.Intro != nil {
		u.Intro = *user.Intro
	}
	return u
}
