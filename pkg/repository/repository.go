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
	Items
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		NewAuthPostgres(db),
		NewItemPostgres(db),
	}
}

type Items interface {
    Create(item model.Item) error
    GetAll(name, sort string) ([]model.Item, error)
    GetById(id int) (model.Item, error)
    Delete(id int) error
    Update(item model.Item, id int) error
	GiveRatingById(rating float32, id int) error 
	FilterbyRating(sort string) ([]model.Item, error)
	FilterbyPrice(sort string) ([]model.Item, error)
}