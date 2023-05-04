package model

type User struct {
	Id       int     `json:"-" db:"id"`
	Username string  `json:"username" binding:"required"`
	Password string  `json:"password" binding:"required"`
	Account  float32 `json:"account" gorm:"default:10000.0"`
}
