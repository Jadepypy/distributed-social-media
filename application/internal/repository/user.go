package repository

import (
	"github.com/Jadepypy/distributed-social-media/application/internal/domain"
	"github.com/Jadepypy/distributed-social-media/application/internal/repository/dao"
	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(user domain.User) error
	GetUserByEmail(email string) (domain.User, error)
	UpdateUser(user domain.User) error
}

type userRepository struct {
	dao dao.UserDAO
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		dao: dao.NewUserDao(db),
	}
}

func (u *userRepository) CreateUser(user domain.User) error {
	//TODO implement me
	panic("implement me")
}

func (u *userRepository) GetUserByEmail(email string) (domain.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u *userRepository) UpdateUser(user domain.User) error {
	//TODO implement me
	panic("implement me")
}
