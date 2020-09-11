package main

import (
	"fmt"
)

func fire(a string) {
	fmt.Println("fire", a)
}
func cloud(a string) {
	fmt.Println("cloud", a)
}
func main() {
	var f func(a string)
	f = fire
	f("ha")
	f = cloud
	f("ha")
}
