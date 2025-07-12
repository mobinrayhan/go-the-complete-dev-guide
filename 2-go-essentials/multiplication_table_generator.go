package main

import "fmt"

const MULTIPLICATION_LENGTH = 10

func main() {
	userInput := 1

	fmt.Print("Enter Your Favorite Number : ")
	_, err := fmt.Scanf("%d", &userInput)

	if err != nil {
		fmt.Println("Input is invalid!")
		panic(err)
	}

	for i := 1; i <= MULTIPLICATION_LENGTH; i++ {
		formattedText := fmt.Sprintf("%d X %d = %d \n", userInput, i, i*userInput)
		fmt.Print(formattedText)
		// fmt.Printf("%d X %d = %d \n", userInput, i, i*userInput)
	}

	// for i := range MULTIPLICATION_LENGTH {
	// 	// formatted =
	// 	fmt.Printf("%d X %d = %d \n", userInput, i, i*userInput)
	// }
}
