package main

import (
	"fmt"

	"mobin.dev/personal_library_manager_cli/models"
)

func main() {

	book := models.Book{
		Title: "Javascript Definition Guide!",
		// Title:         "",
		Author: "Mobin",
		// Author:        "",
		Pages: 200,
		// Pages:         0,
		PublishedYear: 2004,
		IsRead:        true,
		// ShortReview:   "Hello This is short Review",
		ShortReview: "In Go, fmt.Sprintf is a function within the fm data according to a format specifier and returns the resulting Sprintf.",
	}

	newBook, newBookErr := book.New(book)

	if newBookErr != nil {
		fmt.Println(newBookErr)
		panic(newBookErr)
	}

	fmt.Println(newBook)
}
