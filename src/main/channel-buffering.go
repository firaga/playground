package main

import "fmt"

func main() {
	messages := make(chan string, 2)

	go func() {
		messages <- "hello"
		messages <- "world"
		messages <- "!"
	}()

	msg := <-messages
	fmt.Println(msg)
	msg = <-messages
	fmt.Println(msg)
	//msg = <-messages
	//fmt.Println(msg)
}
