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
		Items: NewItemService(repos.Items),
	}
}

type Items interface {
	Create(item model.Item) error
	GetAll(name, sort string) ([]model.Item, error)
	GetById(id int) (model.Item, error)
	Delete(id int) error
	Update(Item model.Item, id int) error
	GiveRatingById(rating float32, id int) error
	FilterbyRating(sort string) ([]model.Item, error)
	FilterbyPrice(sort string) ([]model.Item, error)
}