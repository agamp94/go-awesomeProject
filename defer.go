package main

import "fmt"

func logging() {
	fmt.Println("Logged")
	message := recover()
	if message != nil {
		fmt.Println("error dengan message:", message)
	}
}

func runApp(num rune) {
	defer logging()
	fmt.Println("Starting...")
	test := 10 / num
	fmt.Println(test)

}

func runApp2(error bool) {
	defer logging()
	if error {
		panic("Error")
	}
}

func main() {
	//runApp(10)
	runApp2(true)
	fmt.Println("Agam")

}
