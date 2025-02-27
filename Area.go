package main

import "fmt"

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func main() {
	rect := Rectangle{Width: 5, Height: 10}
	area := rect.Area()
	fmt.Println("Площадь прямоугольника:", area)
}
