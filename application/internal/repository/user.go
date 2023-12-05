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
	//TODO implement me
	panic("implement me")
}

func daoToDomain(user dao.User) domain.User {
	return domain.User{
		Email:    user.Email,
		Password: user.Password,
		Birthday: user.Birthday,
		Nickname: user.Nickname,
		Intro:    user.Intro,
	}
}
