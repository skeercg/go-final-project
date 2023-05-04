package model

type Item struct { 
	Id      int		`json:"-" db:"id"`
	Name  string  	`json:"name" binding:"required"`
	Price  float32 	`json:"price" binding:"required"`
	Rating  float32 `json:"rating"`
}