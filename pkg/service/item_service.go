package service

import (
	"go-final-project/pkg/model"
	"go-final-project/pkg/repository"
)

type ItemService struct {
	repo repository.Items
}

func NewItemService(repo repository.Items) *ItemService {
	return &ItemService{repo: repo}
}

func (r *ItemService) Grade(id int, grade float32) error {
	err := r.repo.Grade(id, grade)

	return err
}

func (r *ItemService) Create(Item model.Item) error {
	err := r.repo.Create(Item)

	return err
}

func (r *ItemService) GetAll(name, sort string) ([]model.Item, error) {
	Items, err := r.repo.GetAll(name, sort)

	return Items, err
}

func (r *ItemService) GetById(id int) (model.Item, error) {
	Item, err := r.repo.GetById(id)

	return Item, err
}

func (r *ItemService) Delete(id int) error {
	err := r.repo.Delete(id)

	return err
}

func (r *ItemService) Update(Item model.Item, id int) error {
	err := r.repo.Update(Item, id)

	return err
}

func (r *ItemService) GiveRatingById(rating float32, id int) error {
	err := r.repo.GiveRatingById(rating, id)

	return err
}

func (r *ItemService) FilterByRating(sort string) ([]model.Item, error) {
	Items, err := r.repo.FilterByRating(sort)

	return Items, err
}

func (r *ItemService) FilterByPrice(sort string) ([]model.Item, error) {
	Items, err := r.repo.FilterByPrice(sort)

	return Items, err
}
