package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"time"
	"github.com/spf13/cobra"
	//"github.com/spf13/cobra"
)

// Book structure
type Book struct {
	Title       string    `json:"title"`
	Author      string    `json:"author"`
	TotalPages  int       `json:"totalPages"`
	CurrentPage int       `json:"currentPage"`
	StartDate   time.Time `json:"startDate"`
}

// Function to add a new book
func addBook(title, author string, totalPages int) {
	data, _ := ioutil.ReadFile("books.json")
	var books []Book
	json.Unmarshal(data, &books)

	newBook := Book{title, author, totalPages, 0, time.Now()}
	books = append(books, newBook)

	updatedData, _ := json.MarshalIndent(books, "", "  ")
	ioutil.WriteFile("books.json", updatedData, 0644)

	fmt.Println("Book added:", title)
}

// Function to update progress
func updateProgress(title string, currentPage int) {
	data, _ := ioutil.ReadFile("books.json")
	var books []Book
	json.Unmarshal(data, &books)

	found := false
	for i, book := range books {
		if book.Title == title {
			books[i].CurrentPage = currentPage
			found = true
			break
		}
	}

	if found {
		updatedData, _ := json.MarshalIndent(books, "", "  ")
		ioutil.WriteFile("books.json", updatedData, 0644)
		fmt.Println("Progress updated for:", title)
	} else {
		fmt.Println("Book not found:", title)
	}
}

// Function to display statistics
func displayStatistics() {
	data, _ := ioutil.ReadFile("books.json")
	var books []Book
	json.Unmarshal(data, &books)

	for _, book := range books {
		daysSpent := int(time.Since(book.StartDate).Hours() / 24)
		readingPace := float64(book.CurrentPage) / float64(daysSpent)

		fmt.Println("Title:", book.Title)
		fmt.Println("Progress:", book.CurrentPage, "/", book.TotalPages)
		fmt.Println("Days Spent:", daysSpent)
		fmt.Printf("Reading Pace: %.1f pages/day\n", readingPace)
		fmt.Println("--------------------")
	}
}

func main() {
	var rootCmd = &cobra.Command{
		Use:   "booktracker",
		Short: "Track your book reading progress",
	}

	var addCmd = &cobra.Command{
		Use:   "add [title] [author] [totalPages]",
		Short: "Add a book",
		Args:  cobra.ExactArgs(3),
		Run: func(cmd *cobra.Command, args []string) {
			title := args[0]
			author := args[1]
			totalPages, _ := strconv.Atoi(args[2])
			addBook(title, author, totalPages)
		},
	}

	var updateCmd = &cobra.Command{
		Use:   "update [title] [currentPage]",
		Short: "Update progress for a book",
		Args:  cobra.ExactArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			title := args[0]
			currentPage, _ := strconv.Atoi(args[1])
			updateProgress(title, currentPage)
		},
	}

	var statsCmd = &cobra.Command{
		Use:   "stats",
		Short: "Display reading statistics",
		Run: func(cmd *cobra.Command, args []string) {
			displayStatistics()
		},
	}

	rootCmd.AddCommand(addCmd, updateCmd, statsCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
