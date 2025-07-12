package main

import (
	"errors"
	"fmt"
)

func main() {

	monthlyIncome, monthlyIncomeErr := getUserInput("Enter Your Monthly Income ")

	if monthlyIncomeErr != nil {
		fmt.Println("Please Enter Monthly Income As Number!")
		panic(monthlyIncomeErr)
	}

	expenses, expensesErr := getUserInput("Enter Your Monthly Income ")

	if expensesErr != nil {
		fmt.Println("Please Enter Expense As Number!")
		panic(expensesErr)
	}

	if monthlyIncome < expenses {
		fmt.Println("Users Spend More Money Then The Income! 😭")
	} else {
		fmt.Printf("You Saved %d TK per Month !", monthlyIncome-expenses)
	}
}

func getUserInput(label string) (int, error) {
	var userInput int

	fmt.Print(label, ": ")
	_, err := fmt.Scanf("%d", &userInput)

	if err != nil {
		return 0, errors.New("input must be a number")
	}

	return userInput, nil
}
