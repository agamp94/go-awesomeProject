package main

import "fmt"

func main() {
	hello := sayHello
	myResp := hello("Agam", "Test")
	fmt.Println(myResp)

	resp2, _ := multipleValFunc("Agamgam", "Test")
	fmt.Println(resp2, "wow")
}

func sayHello(name ...string) string {
	fmt.Println("Hello", name)

	return name[0]
}

func multipleValFunc(names ...string) (name string, count int) {
	return names[0], 1
}
