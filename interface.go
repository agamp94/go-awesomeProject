package main

import "fmt"

type HasName interface {
	GetName() string
}

type Person struct {
	name string
}

func (p Person) GetName() string {
	return p.name
}

func SayHello(hasName HasName) {
	fmt.Println("Hello", hasName.GetName())
}

func Ups(i int) any {
	if i == 1 {
		return 2
	} else if i == 2 {
		return "WOW"
	} else {
		return true
	}
}

func main() {
	orang := Person{
		name: "Agam",
	}

	switch value := Ups(3).(type) {
	case string:
		fmt.Println("String", value)
	case int:
		fmt.Println("Int", value)
	case bool:
		fmt.Println("Boolean", value)
	default:
		fmt.Println("Unknown")
	}
	SayHello(orang)
}
