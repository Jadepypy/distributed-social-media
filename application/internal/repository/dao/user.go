package dao

import (
	"context"
	"database/sql"
	"gorm.io/gorm"
)

type User struct {
	Id       int64          `gorm:"primaryKey,autoIncrement"`
	Email    sql.NullString `gorm:"unique"`
	Password string
	Nickname string
	Intro    string

	// ms, UTF+0:00
	CreatedAt int64
	// ms, UTF+0:00
	UpdatedAt int64
}

func NewUserDao(db *gorm.DB) UserDAO {
	return &userDao{
		db: db,
	}
}

type UserDAO interface {
	GetByEmail(ctx context.Context, email string) (User, error)
	Insert(ctx context.Context, user User) error
	Update(ctx context.Context, user User) error
}

type userDao struct {
	db *gorm.DB
}

func (u *userDao) GetByEmail(ctx context.Context, email string) (User, error) {
	//TODO implement me
	panic("implement me")
}

func (u *userDao) Insert(ctx context.Context, user User) error {
	//TODO implement me
	panic("implement me")
}

func (u *userDao) Update(ctx context.Context, user User) error {
	//TODO implement me
	panic("implement me")
}
