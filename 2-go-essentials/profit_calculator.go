package main

import (
	"fmt"
	"math"
)

func main() {
	var revenue, expenses, taxRate float64

	fmt.Print("Enter the revenue: ")
	fmt.Scan(&revenue)
	fmt.Print("Enter the expenses: ")
	fmt.Scan(&expenses)
	fmt.Print("Enter the tax rate (as a percentage): ")
	fmt.Scan(&taxRate)

	earningBeforeTax := revenue - expenses
	profit := earningBeforeTax - (earningBeforeTax * taxRate / 100)
	ratio := earningBeforeTax / profit

	fmt.Println("Profit Calculator")
	fmt.Println("Earning Before Tax: $", earningBeforeTax)
	fmt.Println("Earning After Tax: $", profit)
	fmt.Println("Profit Ratio: ", math.Round(ratio*100)/100)
	fmt.Println("Just testing math function ", math.Pow(2, 3))
}
