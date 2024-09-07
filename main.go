package main

import (
	"log"

	"go.uber.org/dig"
)

type Animal interface {
	GetName() string
}

type Cat struct {
	name string
}

func (c Cat) GetName() string {
	return "Cat: " + c.name
}

func NewCat() Animal {
	return Cat{
		name: "neko",
	}
}

type Dog struct {
	name string
}

func (c Dog) GetName() string {
	return "Dog: " + c.name
}

func Describe(str string) {
	log.Println(str)
}

func main() {
	c := dig.New()

	err := c.Provide(func(animal Animal) string {
		return animal.GetName()
	})

	if err != nil {
		panic(err)
	}

	if err := c.Provide(NewCat); err != nil {
		panic(err)
	}

	if err := c.Invoke(Describe); err != nil {
		panic(err)
	}
}
