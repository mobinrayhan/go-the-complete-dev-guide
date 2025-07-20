package fsjson

import (
	"encoding/json"
	"fmt"
	"os"

	"mobin.dev/personal_library_manager_cli/models"
)

func WriteToJson(fileName string, jsonData models.Book) {

	convertedJson, err := json.MarshalIndent(jsonData, "", "  ")
	if err != nil {
		fmt.Printf("Error marshaling JSON: %v\n", err)
		return
	}

	err = os.WriteFile(fileName, convertedJson, 0644)

	if err != nil {
		fmt.Printf("Error writing to file: %v\n", err)
		return
	}
}

func ReadFromJson(fileName string) {

}
