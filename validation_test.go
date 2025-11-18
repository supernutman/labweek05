package main

import (
    "testing"
    "myproject/entity"
	
)

// ...existing code...
func TestValidateStudent(t *testing.T) {
    tests := []struct {
        name   string
        s      entity.Student
        wantOK bool
    }{
		{"valid", entity.Student{Name: "John Doe", Price: 100.0, Stock: 10, Description: "A valid description"}, true},
        {"empty name", entity.Student{Name: "", Price: 50.0, Stock: 5, Description: "A valid description"}, false},
        {"negative price", entity.Student{Name: "Jane Doe", Price: -10.0, Stock: 5, Description: "A valid description"}, false},
        {"negative stock", entity.Student{Name: "Alice", Price: 20.0, Stock: -5, Description: "A valid description"}, false},
        {"short desc", entity.Student{Name: "Bob", Price: 30.0, Stock: 15, Description: "Short"}, false},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            ok, err := entity.ValidateStudent(tt.s)
            if ok != tt.wantOK {
                t.Errorf("%s: expected ok=%v got ok=%v err=%v", tt.name, tt.wantOK, ok, err)
            }
        })
    }
}
// ...existing code...