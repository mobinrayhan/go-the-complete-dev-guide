package fileWriter

import (
	"errors"
	"os"
)

func WriteValueToFile(fileName, value string) error {
	err := os.WriteFile(fileName, []byte(value), 0644)

	if err != nil {
		return errors.New("failed to write file")
	}

	return nil
}

func GetValueFromFile(fileName string) (string, error) {
	readData, err := os.ReadFile(fileName)

	if err != nil {
		return "", errors.New("failed to read file")
	}

	return string(readData), nil
}
