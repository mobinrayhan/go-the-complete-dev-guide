package main

import (
	"fmt"

	"mobin.dev/multiplication_table_generator/fileWriter"
)

const MULTIPLICATION_LENGTH = 10
const fileName = "multiplication.txt"

func main() {
	userInput := 1
	fmt.Print("Enter Your Favorite Number : ")
	_, err := fmt.Scanf("%d", &userInput)

	if err != nil {
		fmt.Println("Input is invalid!")
		panic(err)
	}

	finalStr := printMultiplicationTable(userInput)
	existingFileData, err := fileWriter.GetValueFromFile(fileName)

	if err != nil {
		fileWriter.WriteValueToFile(fileName, finalStr)
	} else {
		existingFormattedData := fmt.Sprintf("%s--------------------\n", existingFileData)
		fileWriter.WriteValueToFile(fileName, existingFormattedData+finalStr)
	}
}
