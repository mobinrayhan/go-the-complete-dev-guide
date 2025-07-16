package main

import (
	"fmt"

	bankaccount "mobin.dev/practice/bank-account"
	"mobin.dev/practice/input"
)

func main() {
	var ownerName string
	amount, amountErr := input.GetInputValue("Enter Balance")
	fmt.Print("Enter Your Name : ")
	fmt.Scan(&ownerName)

	if amountErr != nil {
		fmt.Println(amountErr)
		panic(amountErr)
	}

	accUser := bankaccount.BankAccount{
		Owner:   ownerName,
		Balance: &amount,
	}
	fmt.Printf("HI %s, Your Account Balance %.2f\n", ownerName, amount)

	fmt.Println("Chose One Prompt Here : ")
	fmt.Println("1. Withdraw Money")
	fmt.Println("2. Deposit Money")

	promptChose, optionChooseErr := input.GetInputValue("Choosing Options : ")

	if optionChooseErr != nil {
		fmt.Println(optionChooseErr)
		panic(optionChooseErr)
	}

	if promptChose == 1 {
		widthRawAmount, _ := input.GetInputValue("Withdraw ")
		accUser.Withdraw(widthRawAmount)

	} else if promptChose == 2 {
		depositAmount, _ := input.GetInputValue("Deposit ")
		accUser.Deposit(depositAmount)
		// fmt.Println(depositAmount)
	}

	fmt.Println(accUser.MainBalance())
}
