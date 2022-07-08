package main

import (
	"flag"
	"fmt"
	"math"
	"os"
	"strconv"
	"time"
)

func init() {
	fmt.Println("Initialized")
}
func main() {
	// package flag
	args := os.Args
	hostName, err := os.Hostname()
	env := os.Getenv("TEST_PASS")
	fmt.Println(args)
	fmt.Println(env)
	fmt.Println(hostName, err)

	// package flag
	host := flag.String("host", "localhost", "Database host")
	flag.Parse()
	fmt.Println(*host)

	// package strconv
	fmt.Println(strconv.ParseInt("1011", 2, 32))
	fmt.Println(strconv.Atoi("2000"))

	// package math
	fmt.Println(math.Round(123.53))

	// package time
	fmt.Println(time.Now())
	fmt.Println(time.Parse(time.RFC3339, "2022-03-22T15:30:21Z"))
}
