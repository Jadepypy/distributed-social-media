package internal

import (
	"github.com/Jadepypy/distributed-social-media/application/internal/handler"
	"github.com/Jadepypy/distributed-social-media/application/internal/repository"
	"github.com/Jadepypy/distributed-social-media/application/internal/usecase"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewUserHandler() *handler.UserHandler {
	return handler.NewUserHandler(usecase.NewUserUseCase(repository.NewUserRepository(NewDB())))
}

func NewDB() *gorm.DB {
	// TODO: remove hardcoded config
	dsn := "host=localhost user=test password=test dbname=dsm port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db
}
