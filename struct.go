package main

import "fmt"

type Customer struct {
	Name string
	age  byte
}

func (cust Customer) sayHello2(name string) {
	fmt.Println("Hi,", cust.Name, "My name is:", name)
}

func main() {
	var orang Customer

	orang.Name = "Agam"
	orang.age = 28
	orang.sayHello2("Dugong")

	test := Customer{}
	test2 := Customer{"agam3", 33}

	fmt.Println(test)
	fmt.Println(test2)
	fmt.Println(orang)
}
