package main

import "fmt"

type Filter func(string) string

type Blacklist func(string) bool

func spamFilter(word string) string {
	if word == "Anjing" {
		return "..."
	} else {
		return word
	}
}

func sayHelloWithFilter(name string, filter Filter) {
	fmt.Println("Hello", filter(name))
}

func registerUser(name string, blacklist Blacklist) {
	if blacklist(name) {
		fmt.Println("You are blocked", name)
	} else {
		fmt.Println("Hello", name)
	}

}

func main() {

	blacklist := func(name string) bool {
		return name == "Anjing"
	}
	sayHelloWithFilter("Agam", spamFilter)
	sayHelloWithFilter("Anjing", spamFilter)
	registerUser("Anjing", blacklist)
	registerUser("Agam", blacklist)
}
