 package entity

import (
	"github.com/asaskevich/govalidator"
	
)


func init() {	// Custom validator for Price
	govalidator.CustomTypeTagMap.Set("price_range", func(i interface{}, o interface{}) bool {
		v, ok := i.(float64)
		if !ok {
			return false
		}
		return v >= 0
	})

	// Custom validator for Stock
	govalidator.CustomTypeTagMap.Set("stock_range", func(i interface{}, o interface{}) bool {
		v, ok := i.(int)
		if !ok {
			return false
		}
		return v >= 0 
	})


}

func ValidateStudent(s Student) (bool, error) {
    ok, err := govalidator.ValidateStruct(s)
    return ok, err
}

