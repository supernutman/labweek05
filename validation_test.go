package main

import (
    "testing"
    "myproject/entity"
    . "github.com/onsi/gomega"
	
)


func TestStudentValidation(t *testing.T) {

    t.Run("Valid student", func(t *testing.T) {
        g := NewWithT(t) 

        s := entity.Student{
            Name:        "Alice",
            Price:       100,
            Stock:       10,
            Description: "This is a valid product",
        }

       ok, err := entity.ValidateStudent(s)

        g.Expect(err).To(BeNil())
        g.Expect(ok).To(BeTrue())
    })

    t.Run("Name is empty", func(t *testing.T) {
        g := NewWithT(t) 

        s := entity.Student{
            Name:        "",
            Price:       100,
            Stock:       5,
            Description: "valid description",
        }

        ok,err :=  entity.ValidateStudent(s)

        g.Expect(err).ToNot(BeNil())
        g.Expect(err.Error()).To(Equal("Name is required"))
        g.Expect(ok).NotTo(BeTrue())
    })

    t.Run("Price <= 0", func(t *testing.T) {
        g := NewWithT(t)

        s := entity.Student{
            Name:        "Bob",
            Price:       -5,
            Stock:       5,
            Description: "valid description",
        }

        ok,err :=  entity.ValidateStudent(s)

        g.Expect(err).ToNot(BeNil())
        g.Expect(err.Error()).To(Equal("Price must be greater than 0"))
        g.Expect(ok).NotTo(BeTrue())
    })

    t.Run("Stock < 0", func(t *testing.T) {
        g := NewWithT(t)

        s := entity.Student{
            Name:        "Bob",
            Price:       10,
            Stock:       -1,
            Description: "valid description",
        }

        ok,err :=  entity.ValidateStudent(s)

        g.Expect(err).ToNot(BeNil())
        g.Expect(err.Error()).To(Equal("Stock cannot be nagative"))
        g.Expect(ok).NotTo(BeTrue())
   
   
    })

    t.Run("Description too short", func(t *testing.T) {
        g := NewWithT(t)

        s := entity.Student{
            Name:        "Bob",
            Price:       10,
            Stock:       1,
            Description: "short",
        }

        ok,err :=  entity.ValidateStudent(s)
        
        g.Expect(err).NotTo(BeNil())
        g.Expect(err.Error()).To(Equal("Description must be at least 10 characters"))
        g.Expect(ok).NotTo(BeTrue())
    })
}
