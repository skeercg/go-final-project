package repository

import (
	"errors"
	"go-final-project/pkg/model"
	"gorm.io/gorm"
)

type UserPostgres struct {
	db *gorm.DB
}

func NewUserPostgres(db *gorm.DB) *UserPostgres {
	return &UserPostgres{db: db}
}

func (r *UserPostgres) GetAccount(id int) (float32, error) {
	var user model.User

	result := r.db.First(&user, id)

	if result.Error != nil {
		return 0, result.Error
	}

	return user.Account, nil
}

func (r *UserPostgres) Withdraw(id int, amount float32) error {
	var user model.User

	result := r.db.First(&user, id)

	if result.Error != nil {
		return result.Error
	}

	if user.Account < amount {
		return errors.New("not enough money")
	}

	user.Account -= amount

	r.db.Save(user)

	return nil
}
