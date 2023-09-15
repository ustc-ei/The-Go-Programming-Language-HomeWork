package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func GenerateTestData(stringNums int) []string {
	var strs []string
	for i := 0; i < stringNums; i++ {
		var str []rune
		length := rand.Int31n(10)
		for j := 0; j < int(length); j++ {
			char := rand.Int31n(26)
			str = append(str, 'a'+char)
		}
		strs = append(strs, string(str))
	}
	return strs
}

func main() {
	strs := GenerateTestData(1000000)
	start := time.Now()
	strAfterCombination := strs[0]
	for _, s := range strs[1:] {
		strAfterCombination += s
	}
	end := time.Now()
	fmt.Println(end.Sub(start))
	start = time.Now()
	_ = strings.Join(strs, "")
	end = time.Now()
	fmt.Println(end.Sub(start))
}
