package main

import "fmt"

type rect struct {
	weight, height int
}

func (r *rect) area() int {
	return r.weight * r.height
}

func (r *rect) perim() int {
	return 2*r.weight + 2*r.height
}

func main() {
	r := rect{weight: 10, height: 5}
	fmt.Println("area:", r.area())
	fmt.Println("perim:", r.perim())

	rp := &r
	fmt.Println("area:", rp.area())
	fmt.Println("perim:", rp.perim())
}
