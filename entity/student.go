package entity

import (
	"gorm.io/gorm"
	
)
type Student struct {
	gorm.Model
	Name  string  `valid:"required~Name is required"`
	Price float64 `valid:"price_range~Price must be greater than 0"`
	Stock int		`valid:"int,stock_range~Stock cannot be nagative,required"`
	Description string `valid:"required,minstringlength(10)~Description must be at least 10 characters"`
}
