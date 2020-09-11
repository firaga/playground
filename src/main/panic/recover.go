package main

import "fmt"

func f() {
	fmt.Println("Start f")
	defer func() {
		fmt.Println("Deferred in f")
	}()
	g()
	fmt.Println("End f")
}

func g() {
	fmt.Println("Start g")
	defer func() {
		fmt.Println("Deferred in g")
	}()
	h()
	fmt.Println("End g")
}

func h() {
	fmt.Println("Start h")
	defer func() {
		fmt.Println("1st deferred in h")
	}()
	defer func() {
		fmt.Println("2nd deferred in h")
		if p := recover(); p != nil {
			fmt.Printf("Panic found: %v\n", p)
		}
	}()
	defer func() {
		fmt.Println("3rd deferred in h")
		panic("boom 2!")
	}()
	fmt.Println("End h befor boom!")
	panic("boom!")
}

func main() {
	f()
}
