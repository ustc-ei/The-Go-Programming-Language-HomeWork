package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// 存储标准输入中每行字符串中重复的个数
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	/*
		一直循环等待, 直到用户输入回车时
		input.Scan() 才会返回 true or false
	*/
	for input.Scan() {
		counts[input.Text()]++
	}
	/*
		输入完成之后, 统计重复行的个数, 并输入个数大于 1 的行内容
	*/
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
