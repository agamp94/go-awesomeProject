package main

import "fmt"

type Address struct {
	City    string
	Country string
}

func (add *Address) changeCountry(newAddress string) {
	add.Country = newAddress
}

func funcPointer(add *Address) {
	add.Country = "Singapore"
}

func main() {
	add1 := Address{
		City:    "Jkt",
		Country: "Indo",
	}

	add2 := &add1
	*add2 = Address{
		City:    "Kyoto",
		Country: "Jpn",
	}
	funcPointer(&add1)

	add3 := new(Address)

	add2.changeCountry("Wakanda Forever")
	fmt.Println(add1)
	fmt.Println(add2)
	fmt.Println(add3)

}
