package main

import (
	"fmt"
)

type shape interface {
	area() int
	perim() int
}

type rectangle struct {
	width, height int
}

func (r *rectangle) area() int {
	return r.width * r.height
}
func (r rectangle) perim() int {
	return 2*r.width + 2*r.height
}
func main() {
	var newr shape = &rectangle{width: 10, height: 5}
	fmt.Println(newr.area())
	fmt.Println(newr.perim())

	r := rectangle{10, 5}
	fmt.Println(r.area())
	fmt.Println(r.perim())
	//p := &r
	//fmt.Println(p.area())
	//fmt.Println(p.perim())

}
