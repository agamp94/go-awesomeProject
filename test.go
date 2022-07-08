package main

import "fmt"

func main() {
	var name string = "agam"
	var testInt rune = 128
	var name3 = "lumintan"
	name2 := "pamungkas"

	var testInt2 = int16(testInt)
	println(testInt2)
	println(2 * 8)
	var (
		umur1  int8   = 64
		alamat string = "alamatku"
	)

	const myConst string = "My con con"

	fmt.Println(myConst)

	fmt.Println(umur1)
	fmt.Println(alamat)
	fmt.Println(len("Agam"))
	fmt.Println("Agam"[0])
	fmt.Println(string("Agam"[0]))

	fmt.Println(name)
	fmt.Println(name2)
	fmt.Println(name3)
	fmt.Println(testInt)
}
