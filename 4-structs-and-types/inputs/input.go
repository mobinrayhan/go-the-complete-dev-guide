package inputs

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

func GetUserInput(label string) (string, error) {
	fmt.Print(label, " : ")
	reader := bufio.NewReader(os.Stdin)
	userInput, err := reader.ReadString('\n')

	if err != nil {
		return "", errors.New("invalid user input")
	}

	return userInput, err
}
