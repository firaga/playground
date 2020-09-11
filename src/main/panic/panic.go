package main

import "fmt"

func f(ch chan int) {
	defer func() {
		fmt.Println("Deferred by f")
	}()
	g()
	ch <- 0
}

func g() {
	defer func() {
		fmt.Println("Deferred by g")
		panic("3rd explosion!")
	}()
	h()
}

func h() {
	defer func() {
		fmt.Println("Deferred by h")
	}()
	defer func() {
		fmt.Println("Deferred by h2")
		panic("2nd explosion!")
		fmt.Println("Deferred by h2 end")
	}()
	panic("boom!")
}

func main() {
	ch := make(chan int)
	go f(ch)
	<-ch
}
