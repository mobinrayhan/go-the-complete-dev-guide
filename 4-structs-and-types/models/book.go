package models

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"mobin.dev/personal_library_manager_cli/inputs"
)

type Book struct {
	Title         string `json:"name"`
	Author        string `json:"author"`
	Pages         int    `json:"pages"`
	PublishedYear int    `json:"publishedYear"`
	IsRead        bool   `json:"isRead"`
	ShortReview   string `json:"shortReview"`
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

func (b *Book) MarkAsRead() {
	b.IsRead = true
}

func (b Book) Summary() string {
	return b.Title + " is written by " + b.Author + " And its published year is " + strconv.Itoa(b.PublishedYear)
}

func (b *Book) UpdateReview(reviewText string) error {
	if reviewText == "" || len(reviewText) > MAX_LENGTH_OF_REVIEW_TEXT {
		return errors.New("review text cannot be empty or it should be at least " + strconv.Itoa(MAX_LENGTH_OF_REVIEW_TEXT) + " character long")
	}

	b.ShortReview = reviewText
	return nil
}

func BookInputs() (string, string, int, int, string) {
	title, titleErr := inputs.GetUserInput("Enter Book Title")
	PrintError(titleErr)

	author, authorErr := inputs.GetUserInput("Enter Book Name")
	PrintError(authorErr)

	pages, pagesErr := inputs.GetUserInput("Enter Book Pages")
	PrintError(pagesErr)
	removedNewLine := strings.TrimSuffix(pages, "\n")
	pagesNumber, conversionErr := strconv.Atoi(removedNewLine)
	PrintError(conversionErr)

	publishedYear, publishYearErr := inputs.GetUserInput("Enter Book Publish Year")
	PrintError(publishYearErr)
	removedNewLine = strings.TrimSuffix(publishedYear, "\n")
	convertedPublishedYear, convertedPublishYearErr := strconv.Atoi(removedNewLine)
	PrintError(convertedPublishYearErr)

	convertedStr := fmt.Sprintf("Enter Book Short Review Within  %d", MAX_LENGTH_OF_REVIEW_TEXT)
	shortReview, shortReviewErr := inputs.GetUserInput(convertedStr)
	PrintError(shortReviewErr)

	return title, author, pagesNumber, convertedPublishedYear, shortReview
}

func PrintError(err error) {
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

}
