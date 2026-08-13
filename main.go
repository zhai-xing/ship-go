package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	line, _ := r.ReadString('\n')
	line = strings.TrimRight(line, "\r\n")
	// Try to parse and print the result.
	n, err := strconv.Atoi(line)
	if err == nil {
		fmt.Println("ok", n)
		return
	}
	fmt.Println("bad")
}
