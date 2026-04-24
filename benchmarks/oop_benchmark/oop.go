package main

import "fmt"

type Circle struct {
	Radius int
}

func (c *Circle) Area() int {
	return 314 * c.Radius * c.Radius / 100
}

func (c *Circle) Perimeter() int {
	return 628 * c.Radius / 100
}

type Rectangle struct {
	Width  int
	Height int
}

func (r *Rectangle) Area() int {
	return r.Width * r.Height
}

func (r *Rectangle) Perimeter() int {
	return 2 * (r.Width + r.Height)
}

func main() {
	N := 100000
	totalArea := 0
	totalPerimeter := 0

	for i := 0; i < N; i++ {
		r := 1 + i%100

		if i%3 == 0 {
			c := &Circle{Radius: r}
			totalArea += c.Area()
			totalPerimeter += c.Perimeter()
		} else if i%3 == 1 {
			rect := &Rectangle{Width: r, Height: r * 2}
			totalArea += rect.Area()
			totalPerimeter += rect.Perimeter()
		} else {
			sq := &Rectangle{Width: r, Height: r}
			totalArea += sq.Area()
			totalPerimeter += sq.Perimeter()
		}
	}

	fmt.Println(totalArea)
	fmt.Println(totalPerimeter)
}
