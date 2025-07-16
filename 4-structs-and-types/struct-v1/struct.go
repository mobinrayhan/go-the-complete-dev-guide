package main

import (
	"fmt"

	"mobin.dev/practice/area"
	"mobin.dev/practice/input"
)

func main() {
	widthInput, widthErr := input.GetInputValue("Enter The Width Of The Rectangle")
	heightInput, heightErr := input.GetInputValue("Enter The Height Of The Rectangle")

	if widthErr != nil {
		fmt.Println(widthErr)
		panic(widthErr)
	}

	if heightErr != nil {
		fmt.Println(heightErr)
		panic(heightErr)
	}

	rectangle := area.Rectangle{
		Width:  widthInput,
		Height: heightInput,
	}
	areaOfRect := rectangle.Area()

	fmt.Println("Area Of Rectangle : ", areaOfRect)
}
