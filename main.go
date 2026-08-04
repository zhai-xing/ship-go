package main

import (
	"fmt"
)

func main() {
	var num int
	fmt.Scanln(&num)
	fmt.Println(square(num))

}
func square(n int) int {
	return n * n
}
