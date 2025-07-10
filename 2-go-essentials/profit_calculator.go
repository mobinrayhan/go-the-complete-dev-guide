package main

import (
	"fmt"
	"math"
)

func main() {
	var revenue, expenses, taxRate float64

	getPromptWithLabel("Enter the revenue: ", &revenue)
	getPromptWithLabel("Enter the expenses: ", &expenses)
	getPromptWithLabel("Enter the tax rate (as a percentage): ", &taxRate)

	earningBeforeTax := revenue - expenses
	profit := earningBeforeTax - (earningBeforeTax * taxRate / 100)
	ratio := earningBeforeTax / profit

	fmt.Println("Profit Calculator")
	fmt.Println("Earning Before Tax: $", earningBeforeTax)
	fmt.Printf("Earning After Tax: $%.2f\n", profit)
	fmt.Printf("Profit Ratio: %.2f\n", ratio)
	fmt.Println("Just testing math function ", math.Pow(2, 3))
	fmt.Print(greet("Mobin"))
}

func getPromptWithLabel(labelName string, varName any) {
	fmt.Print(labelName)
	fmt.Scan(varName)
}

func greet(name string) (gs, bagGr string) {
	gs = "Hello " + name + " Welcome to golang greet function ! \n"
	bagGr = "Hello " + name + " You are A very bad guy!"
	return gs, bagGr
}
