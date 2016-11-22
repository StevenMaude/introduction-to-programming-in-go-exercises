package main

import (
	"fmt"
	"math"
)

type Shape interface {
	area() float64
	perimeter() float64
}

type Circle struct {
	x, y, r float64
}

type Rectangle struct {
	x1, y1, x2, y2 float64
}

func distance(x1, y1, x2, y2 float64) float64 {
	a := x2 - x1
	b := y2 - y1
	return math.Sqrt(a*a + b*b)
}

func (r *Rectangle) area() float64 {
	return r.length() * r.width()
}

func (r *Rectangle) length() float64 {
	return distance(r.x1, r.y1, r.x1, r.y2)
}

func (r *Rectangle) width() float64 {
	return distance(r.x1, r.y1, r.x2, r.y1)
}

func (r *Rectangle) perimeter() float64 {
	return 2*r.length() + 2*r.width()
}

func (c *Circle) area() float64 {
	return math.Pi * c.r * c.r
}

func (c *Circle) perimeter() float64 {
	return 2 * math.Pi * c.r
}

func totalArea(shapes ...Shape) float64 {
	var area float64
	for _, s := range shapes {
		area += s.area()
	}
	return area
}

func totalPerimeter(shapes ...Shape) float64 {
	var perimeter float64
	for _, s := range shapes {
		perimeter += s.perimeter()
	}
	return perimeter
}

func main() {
	c := Circle{4, 5, 6}
	r := Rectangle{3, 4.4, 7, 10}
	fmt.Println("Circle area", c.area())
	fmt.Println("Rectangle area", r.area())
	fmt.Println("Total area", totalArea(&c, &r))
	fmt.Println("Circle perimeter", c.perimeter())
	fmt.Println("Rectangle perimeter", r.perimeter())
	fmt.Println("Total perimeter", totalPerimeter(&c, &r))
}
