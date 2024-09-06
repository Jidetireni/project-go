package main

import "fmt"

type rect struct {
	width, height int
}

func (r *rect) area() int {
	return r.width * r.height
}

func (r rect) perim() int {
	return 2*r.width + 3*r.height
}

func main() {
	rectangle := rect{
		width:  10,
		height: 5,
	}
	fmt.Println("area: ", rectangle.area())
	fmt.Println("perimeter", rectangle.perim())

	rp := &rect{
		width:  11,
		height: 6,
	}
	fmt.Println("area: ", rp.area())
	fmt.Println("perimeter", rp.perim())
}
