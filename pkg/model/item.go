package model

import "gorm.io/gorm"

type Item struct {
	Id          int     `json:"-" db:"id"`
	Name        string  `json:"name" binding:"required"`
	Price       float32 `json:"price" binding:"required"`
	RatingTotal float32 `json:"-" db:"rating_total"`
	RatingCount int     `json:"-" db:"rating_count"`
	Rating      float32 `json:"rating" gorm:"-"`
}

func (i *Item) AfterFind(tx *gorm.DB) error {
	if i.RatingCount == 0 {
		i.Rating = 0
	} else {
		i.Rating = i.RatingTotal / float32(i.RatingCount)
	}
	return nil
}
