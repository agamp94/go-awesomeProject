package main

import "fmt"

func main() {
	person := map[string]string{
		"name":    "Agam",
		"address": "Assana Vanya Park",
	}

	person["title"] = "Backend Engineer"

	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["title"])
	fmt.Println(len(person))

	var newPerson = make(map[string]string)
	fmt.Println(newPerson)
	fmt.Println(len(newPerson))

	if person["title"] == "Backend Enginee" {
		fmt.Println(person["title"])
	} else {
		fmt.Println("Bukan")
	}

	for count := 1; count <= 10; count++ {
		if count == 6 {
			continue
		}
		fmt.Println("Counter:", count)
	}

	for i, val := range person {
		fmt.Println(i, val)
	}
}
