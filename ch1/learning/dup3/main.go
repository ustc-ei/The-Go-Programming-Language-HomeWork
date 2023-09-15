package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	counts := make(map[string]int)
	for _, fileName := range os.Args[1:] {
		data, err := os.ReadFile(fileName)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup3: %v\n", err)
		}
		strs := strings.Split(string(data), "\n")
		for _, line := range strs {
			if _, ok := counts[line]; !ok {
				fmt.Println(line)
			}
			counts[line]++
		}
	}
}
