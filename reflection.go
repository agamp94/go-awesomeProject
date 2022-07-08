package main

import (
	"fmt"
	"reflect"
)

type Sample struct {
	Name string
}

func main() {
	sample := Sample{"Agam"}
	sampleType := reflect.TypeOf(sample)

	fmt.Println(sampleType)
	fmt.Println(sampleType.Name())
	fmt.Println(sampleType.NumField())
	fmt.Println(sampleType.Field(0))
}
