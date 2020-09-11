package main

import (
	"fmt"
	"strings"
)

type P struct {
	Name string
	Age  int
}

func main() {
	o := []P{
		P{"chain1", 20},
		P{"chain2", 21},
		P{"chain3", 22},
	}
	oPointer := make([]*P, 0, 3)
	newO := make([]P, 0, 3)
	for _, v := range o {
		fmt.Println(v)
		fmt.Println(&v)
		oPointer = append(oPointer, &v)
		newO = append(newO, v)
	}
	fmt.Println(oPointer)
	for _, v := range oPointer {
		fmt.Println(v)
	}
	for _, v := range newO {
		fmt.Println(v)
	}

	data := map[string]string{"1": "A", "2": "B", "3": "C"}
	for k, v := range data {
		data[v] = k
		fmt.Printf("res %v\n", data)
	}

	fmt.Println(strings.Repeat("~", 47))
	data = map[string]string{"1": "A", "2": "B", "3": "C"}
	for i, _ := range data {
		data[i] = "D"
		fmt.Printf("res %v\n", data)
	}
	fmt.Println(strings.Repeat("~", 47))
	data2 := []string{"A", "B", "C"}
	for i, v := range data2 {
		fmt.Println(i, v)
		data2[i] = "D"
		fmt.Printf("res %v\n", data2)
	}

	o = []P{
		P{"chain1", 20},
		P{"chain2", 21},
		P{"chain3", 22},
	}
	for i := 0; i < len(o); i++ {
		o[i] = P{
			Name: o[i].Name,
			Age:  99,
		}
	}
	fmt.Println(o)

	op := []*P{
		&P{"chain1", 20},
		&P{},
		&P{"chain3", 22},
	}
	fmt.Println(strings.Repeat("~", 47))
	fmt.Println(op)
	for _, v := range op {
		fmt.Println(*v)
	}
}
