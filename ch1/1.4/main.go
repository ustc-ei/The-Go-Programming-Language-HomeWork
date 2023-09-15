package main

import (
	"bufio"
	"fmt"
	"os"
)

func ReadTxtFile(fileName string, count map[string]int) {
	f, err := os.Open(fileName)
	if err != nil {
		fmt.Fprintf(os.Stderr, "dup3: %v", err)
		return
	}
	defer f.Close()
	input := bufio.NewScanner(f)
	for input.Scan() {
		if _, ok := count[input.Text()]; ok {
			fmt.Println(f.Name(), input.Text())
		}
		count[input.Text()]++
	}
	fmt.Println(count)
}

func main() {
	count := make(map[string]int)
	for _, fileName := range os.Args[1:] {
		ReadTxtFile(fileName, count)
	}
}
