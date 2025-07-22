package main

import (
	"fmt"

	"mobin.dev/inteface_practice/circle"
	"mobin.dev/inteface_practice/rectangle"
)

type area interface {
	Area() float64
}

func main() {
	newCircle := circle.Circle{
		Radius: 500,
	}

	newRectangle := rectangle.Rectangle{
		Length: 20,
		Width:  30,
	}

	getArea(newCircle)
	getArea(newRectangle)

}

func getArea(a area) {
	returnedArea := a.Area()
	fmt.Println(returnedArea)
}
