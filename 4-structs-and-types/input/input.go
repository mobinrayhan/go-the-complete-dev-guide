package input

import (
	"errors"
	"fmt"
)

// func isValidNumber(input string) bool {
// 	_, err := strconv.ParseFloat(input, 64)
// 	return err == nil
// }

func GetInputValue(label string) (float64, error) {
	var userInput float64

	fmt.Print(label, " : ")
	_, err := fmt.Scan(&userInput)

	if err != nil {
		return 0, errors.New("input is invalid")
	}

	return userInput, nil
}
