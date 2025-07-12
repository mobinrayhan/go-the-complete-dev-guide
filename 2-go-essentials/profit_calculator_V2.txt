package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

const fileName = "calculated_profit.txt"

func getCalculatedResultFromFile(input string) error {
	err := os.WriteFile(fileName, []byte(input), 0644)
	return err
}

func main() {
	revenue := getUserInput("Revenue: ")
	expenses := getUserInput("Expenses: ")
	taxRate := getUserInput("Tax Rate: ")

	// if err != nil {
	// 	fmt.Println(err)
	// 	panic(err)
	// }

	// if err != nil {
	// 	fmt.Println(err)
	// 	panic(err)
	// }

	// if err != nil {
	// 	fmt.Println(err)
	// 	panic(err)
	// }

	ebt, profit, ratio := calculateFinancial(revenue, expenses, taxRate)

	fmt.Printf("%.1f\n", ebt)
	fmt.Printf("%.1f\n", profit)
	fmt.Printf("%.3f\n", ratio)

	// formattedStrEbt := strconv.FormatFloat(ebt, 'f', -1, 64)
	// formattedStrProfit := strconv.FormatFloat(profit, 'f', -1, 64)
	// formattedStrRatio := strconv.FormatFloat(ratio, 'f', -1, 64)

	// finalStr := "Ebt : " + formattedStrEbt + "\n" + "Profit : " + formattedStrProfit + "\n" + "Ratio : " + formattedStrRatio + "\n"
	formattedStr := fmt.Sprintf("EBT : %0.2f\nProfit : %0.2f\nRatio : %0.2f\n", ebt, profit, ratio)
	getCalculatedResultFromFile(formattedStr)
}

func calculateFinancial(revenue, expenses, taxRate float64) (float64, float64, float64) {
	ebt := revenue - expenses
	profit := ebt * (1 - taxRate/100)
	ratio := ebt / profit
	return ebt, profit, ratio
}

func getUserInput(infoText string) float64 {
	var userInput string
	fmt.Print(infoText)
	fmt.Scan(&userInput)

	parsedNumber, err := strconv.ParseFloat(userInput, 64)

	if err != nil {
		err := errors.New("please enter valid number")
		panic(err)
	} else if parsedNumber <= 0 || userInput == "" {
		err := errors.New("number should be positive or at least higher of the zero")
		panic(err)
	} else {
		return parsedNumber
	}
}
