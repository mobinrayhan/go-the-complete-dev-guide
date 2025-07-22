package inputs

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

func GetUserInput(label string) (string, error) {
	fmt.Print(label, " : ")
	reader := bufio.NewReader(os.Stdin)
	userInput, err := reader.ReadString('\n')
	userInput = strings.TrimSuffix(userInput, "\n")

	if err != nil {
		return "", errors.New("invalid user input")
	}

	return userInput, err
}
