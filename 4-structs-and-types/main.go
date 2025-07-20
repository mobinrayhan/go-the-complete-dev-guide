package main

import (
	"fmt"

	"mobin.dev/personal_library_manager_cli/models"
)

func main() {
	title, author, pagesNumber, convertedPublishedYear, shortReview := models.BookInputs()

	bookStruct := models.Book{
		Title:         title,
		Author:        author,
		Pages:         pagesNumber,
		PublishedYear: convertedPublishedYear,
		ShortReview:   shortReview,
		IsRead:        false,
	}

	newBook, err := bookStruct.New(bookStruct)

	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	fmt.Println(newBook)
}
