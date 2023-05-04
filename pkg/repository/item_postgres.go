package repository

import (
	"go-final-project/pkg/model"
	"gorm.io/gorm"
)

type ItemPostgres struct {
	db *gorm.DB
}

func NewItemPostgres(db *gorm.DB) *ItemPostgres {
	return &ItemPostgres{db: db}
}

func (r *ItemPostgres) Grade(id int, grade float32) error {
	var oldItem model.Item

	result := r.db.First(&oldItem, id)

	if result.Error != nil {
		return result.Error
	}

	oldItem.RatingCount += 1
	oldItem.RatingTotal += grade

	result = r.db.Save(&oldItem)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *ItemPostgres) Create(item model.Item) error {
	r.db.Create(&item)

	return nil
}

func (r *ItemPostgres) GetAll(name, sort string) ([]model.Item, error) {
	orderQuery := "price " + sort
	nameQuery := "%" + name + "%"

	var items []model.Item
	result := r.db.Where("name LIKE ?", nameQuery).Order(orderQuery).Find(&items)

	if result.Error != nil {
		return []model.Item{}, result.Error
	}

	return items, nil
}

func (r *ItemPostgres) GetById(id int) (model.Item, error) {
	var item model.Item

	result := r.db.First(&item, id)

	if result.Error != nil {
		return model.Item{}, result.Error
	}

	return item, nil
}

func (r *ItemPostgres) Delete(id int) error {
	result := r.db.Delete(&model.Item{}, id)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *ItemPostgres) Update(item model.Item, id int) error {
	var oldItem model.Item

	result := r.db.First(&oldItem, id)

	if result.Error != nil {
		return result.Error
	}

	oldItem.Name = item.Name
	oldItem.Price = item.Price

	result = r.db.Save(&oldItem)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *ItemPostgres) GiveRatingById(rating float32, id int) error {
	var item model.Item

	result := r.db.First(&item, id)

	if result.Error != nil {
		return result.Error
	}

	item.Rating = rating

	result = r.db.Save(&item)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *ItemPostgres) FilterByRating(sort string) ([]model.Item, error) {
	orderQuery := "rating " + sort

	var items []model.Item
	result := r.db.Order(orderQuery).Find(&items)

	if result.Error != nil {
		return []model.Item{}, result.Error
	}

	return items, nil
}

func (r *ItemPostgres) FilterByPrice(sort string) ([]model.Item, error) {
	orderQuery := "price " + sort

	var items []model.Item
	result := r.db.Order(orderQuery).Find(&items)

	if result.Error != nil {
		return []model.Item{}, result.Error
	}

	return items, nil
}
