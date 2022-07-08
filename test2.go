package main

import "fmt"

func main() {
	// array
	var names = [4]string{
		"agam",
		"pamungkas",
		"lumintan",
	}

	names[3] = "test"
	fmt.Println(names)
	fmt.Println(len(names))

	var mySlice = names[1:3]
	fmt.Println(mySlice)
	fmt.Println(cap(mySlice))
	fmt.Println(len(mySlice))

	names[1] = "gamgam"
	fmt.Println(mySlice)
	fmt.Println(names)

	mySlice2 := append(mySlice, "Ups")
	mySlice2 = append(mySlice2, "Upsie")
	fmt.Println(mySlice)
	fmt.Println(mySlice2)
	fmt.Println(names)

	newSlice := make([]string, 2, 5)
	newSlice = append(newSlice, "wowowo")
	fmt.Println(newSlice)
	fmt.Println(newSlice[0])
	fmt.Println(len(newSlice))

	copySlice := make([]string, len(newSlice), cap(newSlice))
	copy(copySlice, newSlice)
	fmt.Println(copySlice)

	iniSlice := []int{1, 2, 3, 4, 5}
	iniSlice2 := make([]int, 5, 5)
	copy(iniSlice2, iniSlice)
	iniSlice2[1] = 10
	fmt.Println(iniSlice)
	fmt.Println(iniSlice2)
	fmt.Println(len(iniSlice))
	fmt.Println(cap(iniSlice))

}
