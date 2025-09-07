package main

import "fmt"

func sayHelloWorld(name string) {
	fmt.Printf("Hello %s\n", name)
}

func evenOrOd(input int) {
	if (input % 2) == 1 {
		fmt.Println("Its Odd Number")
	} else {
		fmt.Println("Its Even Number")
	}
}

func sum(nums []int) int {
	var sumValue int = 0
	for _, v := range nums {
		sumValue += v
	}
	return sumValue
}

func reverseString(input string) string {
	// V1
	var revString string = ""

	// for _, v := range input {
	// 	value := fmt.Sprintf("%c", v)
	// 	revString = value + revString
	// }

	// return revString

	for i := 0; i < len(input); i++ {
		value := fmt.Sprintf("%c", input[i])
		revString = value + revString
	}
	return revString
}

func factorial(input int) int {
	var factTotal int = 1
	for i := 1; i < input; i++ {
		factTotal *= input
	}
	return factTotal
}

func main() {
	sayHelloWorld("Mobin")
	evenOrOd(545)

	array := []int{10, 20, 30, 40, 50}

	totalSum := sum(array)
	fmt.Println(totalSum)

	revString := reverseString("Hello World")
	fmt.Println(revString)

	fmt.Println(factorial)
}
