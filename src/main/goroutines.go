package main

import "fmt"

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}




func main() {
	f("hello")
	go f("routines")

	msg:="heep"
	go func() {
		fmt.Println(msg)
	}()

	fmt.Scanln()
	fmt.Print("done")
}
