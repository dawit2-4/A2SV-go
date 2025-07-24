package controllers

import (
	"Library-management/models"
	"Library-management/services"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type LibraryController struct {
	Manager services.LibraryManager
}

func (c *LibraryController) Run() {
	scanner := bufio.NewScanner(os.Stdin)

for {
		fmt.Println("\nLibrary Management System")
		fmt.Println("1. Add Book")
		fmt.Println("2. Remove Book")
		fmt.Println("3. Borrow Book")
		fmt.Println("4. Return Book")
		fmt.Println("5. List Available Books")
		fmt.Println("6. List Borrowed Books")
		fmt.Println("7. Add Member")
		fmt.Println("8. Exit")
		fmt.Print("Choose an option: ")

		scanner.Scan()
		choice , err := strconv.Atoi(strings.TrimSpace(scanner.Text()))

		if err != nil {
			fmt.Println("Invalid input, please enter a number")
			continue
		}
		switch choice {
		case 1:
				c.addBook(scanner)
		case 2:
				c.removeBook(scanner)
		case 3:
				c.borrowBook(scanner)
		case 4:
				c.returnBook(scanner)
		case 5:
				c.listAvailableBooks()
		case 6:
				c.listBorrowedBooks(scanner)
		case 7:
				c.addMember(scanner)
		case 8:
				fmt.Println("Exiting...")
				return
		default:
				fmt.Println("Invalid option")
		}
	}
}

func (c *LibraryController) addBook(scanner *bufio.Scanner){
	fmt.Println("Enter Book Title:")
	scanner.Scan()
	title := strings.TrimSpace(scanner.Text())
	fmt.Println("Enter the Author:")
	scanner.Scan()
	author := strings.TrimSpace(scanner.Text())

	book := models.Book{
		Title: title,
		Author: author,
	}

	bookID := c.Manager.AddBook(book)
	fmt.Printf("Book Added Succesfully with ID %d \n", bookID)
}

func (c *LibraryController) removeBook(scanner *bufio.Scanner){
	fmt.Println("Enter book ID:")
	scanner.Scan()
	id, err := strconv.Atoi(strings.TrimSpace(scanner.Text()))

	if err != nil {
		fmt.Println("Enter a vaild ID")
		return
	}
	c.Manager.RemoveBook(id)
	fmt.Println("Book removed successfully")
}

func (c *LibraryController) borrowBook(scanner *bufio.Scanner) {
	fmt.Println("Enter book ID:")
	scanner.Scan()
	bookID, err := strconv.Atoi(strings.TrimSpace(scanner.Text()))
	if err != nil{
		fmt.Println("Enter a valid ID")
	}
	fmt.Println("Enter member ID:")
	scanner.Scan()
	memberID, err := strconv.Atoi(strings.TrimSpace(scanner.Text()))
	if err != nil{
		fmt.Println("Enter a vaild memberID")
	}
	err = c.Manager.BorrowBook(bookID, memberID)
	if err != nil{
		fmt.Printf("%v\n",err)
		return
	}
	fmt.Println("Book Borrowed successfully")
}

func (c *LibraryController) returnBook(scanner *bufio.Scanner){
	fmt.Println("Enter the book ID:")
	scanner.Scan()
	bookID, err := strconv.Atoi(strings.TrimSpace(scanner.Text()))
	if err != nil{
		fmt.Println("Enter a valid bookID")
	}
	fmt.Println("Enter member ID:")
	scanner.Scan()
	memberID, err := strconv.Atoi(strings.TrimSpace(scanner.Text()))
	if err != nil {
		fmt.Println("Enter a valid member ID")
	}
	err = c.Manager.ReturnBook(bookID, memberID)
	if err != nil{
		fmt.Printf("%v\n", err)
	}
	fmt.Println("Book returned successfully")
}

func (c*LibraryController) listAvailableBooks() {
	books := c.Manager.ListAvailableBooks()
	if len(books) == 0{
		fmt.Println("No Available Book")
		return
	}
	fmt.Println("Available books:")
	for _, book := range books{
		fmt.Printf("ID: %d, Title: %s, Author:%s", book.ID, book.Title, book.Author)
	}
}

func (c *LibraryController) listBorrowedBooks(scanner *bufio.Scanner) {
    fmt.Print("Enter Member ID: ")
    scanner.Scan()
    memberID, err := strconv.Atoi(strings.TrimSpace(scanner.Text()))
    if err != nil {
        fmt.Println("Invalid Member ID")
        return
    }
    books := c.Manager.ListBorrowedBooks(memberID)
    if len(books) == 0 {
        fmt.Println("No borrowed books")
        return
    }
    fmt.Println("Borrowed Books:")
    for _, book := range books {
        fmt.Printf("ID: %d, Title: %s, Author: %s\n", book.ID, book.Title, book.Author)
    }
}

func (c *LibraryController) addMember(scanner *bufio.Scanner) {
    fmt.Print("Enter Member Name: ")
    scanner.Scan()
    name := strings.TrimSpace(scanner.Text())

    member := models.Member{
        Name: name,
    }
    memberID := c.Manager.AddMember(member)
    fmt.Printf("Member added successfully with ID %d\n", memberID)
}