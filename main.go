package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	line, _ := r.ReadString('\n')
	words := strings.Fields(strings.TrimSpace(line))
	seen := map[string]int{}
	// Add each word and print the size of seen.
	for _, word := range words {
		seen[word]++
	}
	var sum = len(seen)
	fmt.Println(sum)
}
