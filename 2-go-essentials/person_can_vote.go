package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {

	name := getInput("Enter Your Name: ")
	ageStr := getInput("Enter Your Age: ")

	ageStrToInt, err := strconv.Atoi(ageStr)

	println("Data Type : ", reflect.TypeOf(ageStrToInt).Kind())
	println(err)

	if err != nil {
		fmt.Println("Please Enter Age As Number", err)
		return
	}

	if !canVote(ageStrToInt) {
		fmt.Println("You Are Not Eligible For Voting, ", name, "!")
	} else {
		fmt.Printf("%s is eligible to vote. \n", name)
	}
}

func getInput(label string) string {
	var userInput string

	fmt.Print(label)
	fmt.Scan(&userInput)

	return userInput
}

func canVote(age int) bool {
	return age >= 18
}
