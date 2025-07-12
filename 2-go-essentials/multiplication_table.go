package main

import "fmt"

func printMultiplicationTable(n int) string {
	printStr := ""

	for i := 1; i <= MULTIPLICATION_LENGTH; i++ {
		formattedText := fmt.Sprintf("%d X %d = %d \n", n, i, i*n)
		fmt.Print(formattedText)
		printStr += formattedText
		// fmt.Printf("%d X %d = %d \n", userInput, i, i*userInput)
	}

	return printStr
}
