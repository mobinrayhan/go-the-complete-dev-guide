package main

import "fmt"

const TOTAL_SUBJECT = 3

func main() {

	calculatedValue := 0

	// for i := 0; i < 3; i++ {
	for i := range TOTAL_SUBJECT {
		var userInput int
		fmt.Print("Enter Your ", i+1, " Subject Marks : ")
		_, err := fmt.Scanf("%d", &userInput)

		if err != nil {
			panic(err)
		}

		if userInput < 0 || userInput > 100 {
			panic("User input Must be Positive and Under 100")
		}

		calculatedValue += userInput
	}

	avgValue := calculatedValue / TOTAL_SUBJECT

	switch {
	case 80 <= avgValue:
		fmt.Println("You Got A")
	case 70 <= avgValue:
		fmt.Println("You Got B")
	case 60 <= avgValue:
		fmt.Println("You Got C")
	default:
		fmt.Println("You Are Failed! Try Again!")
	}
}
