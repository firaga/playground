package main

import "fmt"

type I interface {
	Get() int
	Set(int)
}

type S struct {
	Age int
}

func (s S) Get() int {
	return s.Age
}

func (s *S) Set(age int) {
	s.Age = age
}

func main() {
	var s I = &S{}
	fmt.Printf("%d\n", s.Get())
	s.Set(20)
	fmt.Printf("%d\n", s.Get())
}
