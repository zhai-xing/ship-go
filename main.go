package main

import (
	"fmt"
)

func main() {
	var num int
	fmt.Scanln(&num)
	var sum int
	for i := 1; i <= num; i++ {
		sum += i
	}
	fmt.Println(sum)
}
