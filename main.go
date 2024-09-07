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

func NewCat1() Cat {
	return Cat{
		name: "neko 2",
	}
}

func NewCat2() Cat {
	return Cat{
		name: "neko 2",
	}
}

type Dog struct {
	name string
}

func NewDog() Dog {
	return Dog{
		name: "inu",
	}
}

func (c Dog) GetName() string {
	return "Dog: " + c.name
}

func Describe(str string) {
	log.Println(str)
}

func SimpleDemo() {
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

	// cannot provide function "main".NewCat (/home/user/dig-test/main.go:21): cannot provide main.Animal from [0]: already provided by "main".NewCat (/home/user/dig-test/main.go:21)
	// if err := c.Provide(NewCat); err != nil {
	// 	panic(err)
	// }

	if err := c.Invoke(Describe); err != nil {
		panic(err)
	}
}

type DogAndCat struct {
	dig.In

	Dog Dog
	Cat Cat
}

func ParameterGroupDemo() {
	c := dig.New()
	if err := c.Provide(NewCat1); err != nil {
		panic(err)
	}
	if err := c.Provide(NewDog); err != nil {
		panic(err)
	}

	err := c.Invoke(func(dogAndCat DogAndCat) {
		log.Println(dogAndCat.Cat.GetName())
		log.Println(dogAndCat.Dog.GetName())
	})
	if err != nil {
		panic(err)
	}
}

func main() {
	SimpleDemo()
	ParameterGroupDemo()
}
