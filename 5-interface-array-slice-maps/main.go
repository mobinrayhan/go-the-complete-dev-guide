package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

func main() {

	fmt.Println("Max of 10 and 20:", Max(10, 20))                                            // int
	fmt.Println("Max of 3.14 and 6.28:", Max(3.14, 6.28))                                    // float64
	fmt.Println("Max of 'apple' and 'banana':", Max("apple adfasdfa afdasd asdf", "banana")) // string
}

func Max[T constraints.Ordered](a, b T) T {
	if a > b {
		return a
	}

	return b
}
