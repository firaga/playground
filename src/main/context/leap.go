package main

import (
	"fmt"
	"time"
)

func genLeak() <-chan int {
	ch := make(chan int)
	go func() {
		var n int
		for {
			ch <- n
			n++
			time.Sleep(time.Second)
		}
	}()
	return ch
}

func main() {
	for n := range genLeak() {
		fmt.Println(n)
		if n == 5 {
			break
		}
	}
	// ……
}