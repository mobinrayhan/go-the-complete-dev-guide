package main

import (
	"fmt"
	"math"
)

func main() {
	a := 20
	b := 20
	result := 0.0
	// swap(&a, &b)

	// println("a : ", a, " b : ", b)
	// randNumber := 20
	// println(randNumber)
	// doubleValue(&randNumber)
	// println(randNumber)

	dividedResult := safeDivide(a, b, &result)
	fmt.Println(result)
	fmt.Println(dividedResult)
	fmt.Println(float64(10000000 / 20))
}

func doubleValue(num *int) {
	value := math.Pow(float64(*num), 2)
	*num = int(value)
}

func swap(a, b *int) {
	tempA := *a
	// temp := *b

	*a = *b
	*b = tempA
}

func safeDivide(a, b int, result *float64) bool {
	*result = float64(a) / float64(b)
	return *result <= 0
}
