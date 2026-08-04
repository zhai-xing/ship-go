package main

import (
	"fmt"
)

func main() {
	var name string
	var age int
	fmt.Scanln(&name)
	fmt.Scanln(&age)
	fmt.Printf("Hi, %s! You are %d years old.\n", name, age)
}
