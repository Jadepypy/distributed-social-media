package dao

import (
	"context"
	"gorm.io/gorm"
	"time"
)

// TODO: use embedding struct for user profile
type User struct {
	Id       int64   `gorm:"primaryKey,autoIncrement"`
	Email    string  `gorm:"unique,not null"`
	Password string  `gorm:"not null"`
	Birthday *string `gorm:"default:null"`
	Nickname *string
	Intro    *string

	// ms, UTF+0:00
	CreatedAt int64
	// ms, UTF+0:00
	UpdatedAt int64 `gorm:"autoUpdateTime:milli"`
}

func NewUserDao(db *gorm.DB) UserDAO {
	return &userDao{
		db: db,
	}
}

type UserDAO interface {
	GetByEmail(ctx context.Context, email string) (User, error)
	Insert(ctx context.Context, user User) error
	Update(ctx context.Context, email string, user User) error
}

type userDao struct {
	db *gorm.DB
}

func (u *userDao) GetByEmail(ctx context.Context, email string) (User, error) {
	var user User
	if err := u.db.Where("email = ?", email).First(&user).Error; err != nil {
		return User{}, err
	}
	return user, nil
}

func (u *userDao) Insert(ctx context.Context, user User) error {
	now := time.Now().UnixMilli()
	user.CreatedAt = now
	user.UpdatedAt = now

	// TODO: customize email unique constraint error
	return u.db.Create(&user).Error
}

func (u *userDao) Update(ctx context.Context, email string, user User) error {
	return u.db.Model(&User{}).Where("email = ?", email).Updates(user).Error
}
