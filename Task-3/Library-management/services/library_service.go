package services

import (
	"Library-management/models"
	"errors"
	"fmt"
)

var (
	ErrBookNotFound = errors.New("book not found")
	ErrBookAlreadyBorrowed = errors.New("book is already borrowed")
	ErrMemberNotFound = errors.New("member not found")
	ErrBookNotBorrowed = errors.New("book was not borrowed by this member")
)

type LibraryManager interface {
    AddBook(book models.Book) int // Return the assigned book ID
    AddMember(member models.Member) int // Return the assigned member ID
    RemoveBook(bookID int)
    BorrowBook(bookID int, memberID int) error
    ReturnBook(bookID int, memberID int) error
    ListAvailableBooks() []models.Book
    ListBorrowedBooks(memberID int) []models.Book
}


type Library struct {
		Books map[int]models.Book
		Members map[int]models.Member
		nextBookId int
		nextMemberId int
}

func NewLibrary() *Library {
	return &Library{
		Books: make(map[int]models.Book),
		Members: make(map[int]models.Member),
		nextBookId: 1,
		nextMemberId: 1,
	}
}

func (l *Library) AddBook(book models.Book) int {
	book.ID = l.nextBookId
	book.Status = models.Available
	l.Books[book.ID] = book
	l.nextBookId ++
	return book.ID
}

func (l *Library) AddMember(member models.Member) int {
	member.ID = l.nextMemberId
	member.BorrowedBooks = make([]models.Book, 0)
	l.Members[member.ID] = member
	l.nextMemberId ++
	return member.ID
}

func (l *Library) RemoveBook(bookID int) {
    if _, exists := l.Books[bookID]; !exists {
        fmt.Printf("Error: %v\n", ErrBookNotFound)
        return
    }
    delete(l.Books, bookID)
}

func (l *Library) BorrowBook(bookID int, memberID int) error {
	book, bookExists := l.Books[bookID] 
	member, memberExists := l.Members[memberID]

	if !bookExists {
		return fmt.Errorf("%w: %d", ErrBookNotFound, bookID)
	}
	if !memberExists {
		return fmt.Errorf("%w: %d", ErrMemberNotFound, bookID)
	}

	if book.Status == models.Borrowed {
		return fmt.Errorf("%w: %d", ErrBookAlreadyBorrowed, bookID)
	}

	book.Status = models.Borrowed
	l.Books[bookID] = book
	member.BorrowedBooks = append(member.BorrowedBooks, book)
	l.Members[memberID] = member
	return  nil
}

func (l *Library) ReturnBook(bookID int, memberID int) error {
	book, bookExists := l.Books[bookID]
	member, memberExists := l.Members[memberID]

	if !bookExists {
		return fmt.Errorf("%w: %d", ErrBookNotFound, bookID)
	}
	if !memberExists {
		return fmt.Errorf("%w: %d", ErrMemberNotFound, memberID)
	}

	if book.Status == models.Available {
		return fmt.Errorf("%w, %d", ErrBookNotBorrowed, bookID)
	}

	for i, b := range member.BorrowedBooks {
		if b.ID == bookID {
			member.BorrowedBooks = append(member.BorrowedBooks[i:], member.BorrowedBooks[i+1:]...)
			book.Status = models.Available
			l.Books[bookID] = book
			l.Members[memberID] = member
			return nil
		}
	}
	return fmt.Errorf("%w book %d member %d", ErrBookNotBorrowed, bookID, memberID)
}

func (l *Library) ListAvailableBooks() []models.Book {
	var available []models.Book
	for _, book := range l.Books {
		if book.Status == models.Available {
			available = append(available, book)
		}
	}
	return available
}

func (l *Library) ListBorrowedBooks(memberID int) []models.Book {
	member, exists := l.Members[memberID]
	if !exists {
		fmt.Printf("%v \n", ErrMemberNotFound)
		return nil
	}
	return member.BorrowedBooks
}