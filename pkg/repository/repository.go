package repository

import (
	"go-final-project/pkg/model"
	"gorm.io/gorm"
)

type Authorization interface {
	CreateUser(user model.User) error
	GetUser(username, password string) (*model.User, error)
}

type Items interface {
	Grade(id int, grade float32) error
	Create(item model.Item) error
	GetAll(name, sort string) ([]model.Item, error)
	GetById(id int) (model.Item, error)
	Delete(id int) error
	Update(item model.Item, id int) error
	GiveRatingById(rating float32, id int) error
	FilterByRating(sort string) ([]model.Item, error)
	FilterByPrice(sort string) ([]model.Item, error)
}

type Users interface {
	GetAccount(id int) (float32, error)
	Withdraw(id int, amount float32) error
}

type Repository struct {
	Authorization
	Items
	Users
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		NewAuthPostgres(db),
		NewItemPostgres(db),
		NewUserPostgres(db),
	}
}
