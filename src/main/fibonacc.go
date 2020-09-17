package main

import "fmt"

//go:nosplit
func fibonacciAwesome(n int64) int64 // 声明但不实现

func main() {
	//fmt.Println(fibonacciDumb(50))
	fmt.Println(fibonacciAwesome(50))
	return
}
