package main

import (
	"fmt"
	"math"
)

type Circle struct {
	Radius float32
}

func (c Circle) Area() float32 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float32 {
	return 2 * math.Pi * c.Radius
}

type Rectangle struct {
	Height float32
	Width  float32
}

func (r Rectangle) Area() float32 {
	return r.Height * r.Width
}

func (r Rectangle) Perimeter() float32 {
	return 2 * (r.Height + r.Width)
}

/*
func PrintCircleArea(c Circle) {
	fmt.Printf("Area = %v\n", c.Area())
}

func PrintRectangleArea(r Rectangle) {
	fmt.Printf("Area = %v\n", r.Area())
}
*/

type ShapeWithArea interface {
	Area() float32
}

func PrintArea(sa ShapeWithArea) {
	fmt.Printf("Area = %v\n", sa.Area())
}

type ShapeWithPerimeter interface {
	Perimeter() float32
}

func PrintPerimeter(sp ShapeWithPerimeter) {
	fmt.Printf("Perimeter = %v\n", sp.Perimeter())
}

type ShapeWithStats interface {
	ShapeWithArea
	ShapeWithPerimeter
}

func PrintStats(ss ShapeWithStats) {
	PrintArea(ss)
	PrintPerimeter(ss)
}

func main() {
	c := Circle{Radius: 10}
	//fmt.Printf("Area = %v\n", c.Area())
	//PrintCircleArea(c)
	/*
		PrintArea(c)
		PrintPerimeter(c)
	*/
	PrintStats(c)
	r := Rectangle{Height: 10, Width: 12}
	//fmt.Printf("Area = %v\n", r.Area())
	//PrintRectangleArea(r)
	/*
		PrintArea(r)
		PrintPerimeter(r)
	*/
	PrintStats(r)
}
