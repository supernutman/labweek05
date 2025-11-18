package main

import (
    "fmt"
    // "github.com/asaskevich/govalidator"
    "myproject/entity"
)

func main() {
    student1 := entity.Student{
        Name:  "eee",
        Price: 0,
        Stock: 0,
        Description: "A ",

    }

    ok, err := entity.ValidateStudent(student1)
    fmt.Println("Student 1 Validation:")
    fmt.Println("OK:", ok)
    fmt.Println("Error:", err)

// type User	struct {
//   Name string      `valid:"required~Name is required"`
//   Age  float64         `valid:"re~age must more than 0"`
//   Stock int         `valid:"int,range(0|1000)~Stock cannot be nagative"`
//   Meta string       `valid:"minstringlength(10)~Description must be at least 10 characters"`
// }

// func init() {
// 	// Custom validator for Age
// 	govalidator.CustomTypeTagMap.Set("age_range", func(i interface{}, o interface{}) bool {
// 		v, ok := i.(float64)
// 		if !ok {
// 			return false
// 		}
// 		return v >= 0 && v <= 10000
// 	})

// 	// Custom validator for Stock
// 	govalidator.CustomTypeTagMap.Set("stock_range", func(i interface{}, o interface{}) bool {
// 		v, ok := i.(int)
// 		if !ok {
// 			return false
// 		}
// 		return v >= 0 && v <= 1000
// 	})
// }

// result, err := govalidator.ValidateStruct(User{"ser", 20, -20, "meta"})
// if err != nil {
// 	println("error: " + err.Error())
// }
// println(result,err)
   
}