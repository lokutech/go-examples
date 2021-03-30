package main

import (
	"fmt"
	"reflect"
)

type Doctor struct {
	number     int
	actorName  string
	companions []string
	episodes   []string // It is ok to declare and not use a struct value
}

type Doctor2 struct {
	number     int
	actorName  string
	companions []string
}

type Animal struct {
	Name   string
	Origin string
}

type Bird struct {
	Animal
	SpeedKPH float32
	CanFly   bool
}

type Person struct {
	Name string `required max:"100"`
	Age  int
}

func main() {
	// STRUCTS
	aDoctor := Doctor{ // Literal syntax
		number:    3,
		actorName: "John Pertwee",
		companions: []string{
			"Liz Shaw",
			"Jo Grant",
		},
	}

	aDoctor2 := Doctor2{ // Positional syntax. It is discouraged because confusing
		3,
		"John Pertwee",
		[]string{
			"Liz Shaw",
			"Jo Grant",
		},
	}

	aDoctor3 := struct{ name string }{name: "John Pertwee"} // An anonymous struct. First curly defines structure of the struct, seconds curly initializer which provides data into struct
	// aDoctor4 := aDoctor3  Make an independent copy, aDoctor4 doesn't mutate aDoctor3
	// aDoctor4 := &aDoctor3  Points to the underlying data, mutates

	fmt.Println(aDoctor)
	fmt.Println(aDoctor2)
	fmt.Println(aDoctor3)
	fmt.Println(aDoctor.number)
	fmt.Println(aDoctor.companions)
	fmt.Println(aDoctor.companions[1])

	// EMBEDDING STRUCTS
	b := Bird{}
	b.Name = "Emu"
	b.Origin = "Australia"
	b.SpeedKPH = 48
	b.CanFly = false

	// with literal syntax
	c := Bird{
		Animal:   Animal{Name: "Emu", Origin: "Australia"},
		SpeedKPH: 48,
		CanFly:   false,
	}
	fmt.Println(b)
	fmt.Println(b.Name)
	fmt.Println(c)

	// TAGS
	t := reflect.TypeOf(Person{})
	field, _ := t.FieldByName("Name")
	fmt.Println(field.Tag)
}
