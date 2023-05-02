package repository

import (
	"go-final-project/pkg/model"

	"gorm.io/gorm"
)

type Authorization interface {
	CreateUser(user model.User) error
	GetUser(username, password string) (*model.User, error)
}

type Repository struct {
	Authorization
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		NewAuthPostgres(db),
	}
}
