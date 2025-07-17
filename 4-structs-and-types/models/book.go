package models

import (
	"errors"
	"fmt"
)

type Book struct {
	Title         string
	Author        string
	Pages         int
	PublishedYear int
	IsRead        bool
	ShortReview   string
}

const MAX_LENGTH_OF_REVIEW_TEXT = 200

func (b Book) New(book Book) (Book, error) {
	if book.Title == "" {
		return Book{}, errors.New("title cannot be empty")
	} else if book.Author == "" {
		return Book{}, errors.New("author cannot be empty")
	} else if book.Pages <= 0 {
		return Book{}, errors.New("page must be positive")
	} else if len(book.ShortReview) > MAX_LENGTH_OF_REVIEW_TEXT {
		formattedText := fmt.Sprintf("Reviews Maximum Length Is %d", MAX_LENGTH_OF_REVIEW_TEXT)
		return Book{}, errors.New(formattedText)
	}
	return book, nil
}
