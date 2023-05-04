package service

import (
	"go-final-project/pkg/model"
	"go-final-project/pkg/repository"
)

type AuthService interface {
	CreateUser(user model.User) error
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type Service struct {
	AuthService
	Items
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		AuthService: NewAuthService(repos.Authorization),
		Items:       NewItemService(repos.Items),
	}
}

type Items interface {
	Grade(id int, grade float32) error
	Create(item model.Item) error
	GetAll(name, sort string) ([]model.Item, error)
	GetById(id int) (model.Item, error)
	Delete(id int) error
	Update(Item model.Item, id int) error
	GiveRatingById(rating float32, id int) error
	FilterByRating(sort string) ([]model.Item, error)
	FilterByPrice(sort string) ([]model.Item, error)
}
