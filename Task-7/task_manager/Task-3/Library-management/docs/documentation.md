# 📚 Library Management System (Console App in Go)

## 📜 Overview

This is a simple, console-based **Library Management System** written in **Go**. The application allows users to manage members and books, and handle book borrowing and returning. It demonstrates practical software engineering principles like **clean architecture**, **interface-based design**, and **state management using pointers**. Ideal for learning struct organization, dependency injection, and Go interfaces.

---

## 📁 Folder Structure

```
library_management/
├── main.go                     # Entry point of the application
├── go.mod                      # Go module file
├── controllers/
│   └── library_controller.go   # CLI interface and user input handling
├── models/
│   ├── book.go                 # Book struct and status definitions
│   └── member.go               # Member struct
├── services/
│   ├── library_service.go      # Core business logic
│   └── library_service_test.go # Unit tests for service layer
├── docs/
│   └── documentation.md        # This documentation file
```

---

## 🧱 Core Architecture

### ✅ Separation of Concerns

- **Models** define data (Book, Member).
- **Services** contain business logic.
- **Controllers** handle user input and console interaction.

### ✅ Interface-Based Design

The `LibraryManager` interface defines all library operations, allowing the controller to work with **any implementation** (mockable, testable, swappable).

### ✅ Dependency Injection

The `LibraryController` receives a `LibraryManager` instance through its constructor, promoting decoupling and testability.

### ✅ Pointer-Based State Management

The system uses **maps of pointers** (`map[int]*Book`) instead of storing copies. This means:

- Changes are reflected immediately across the system
- No need to "write back" updated structs
- Memory-efficient and bug-resistant

---

## 📆 Models

### 📘 `Book` (models/book.go)

| Field  | Type     | Description                |
| ------ | -------- | -------------------------- |
| ID     | `int`    | Auto-incremented unique ID |
| Title  | `string` | Title of the book          |
| Author | `string` | Author of the book         |
| Status | `string` | "Available" or "Borrowed"  |

### 👤 `Member` (models/member.go)

| Field         | Type      | Description                |
| ------------- | --------- | -------------------------- |
| ID            | `int`     | Auto-incremented unique ID |
| Name          | `string`  | Name of the member         |
| BorrowedBooks | `[]*Book` | Slice of pointers to books |

---

## 🤠 Interface

### `LibraryManager` (services/library_service.go)

```go
type LibraryManager interface {
	AddBook(book Book) int
	AddMember(member Member) int
	RemoveBook(bookID int)
	BorrowBook(bookID int, memberID int) error
	ReturnBook(bookID int, memberID int) error
	ListAvailableBooks() []Book
	ListBorrowedBooks(memberID int) []Book
}
```

---

## ⚙️ Implementation

### 🏩 `Library` Struct

The `Library` struct implements `LibraryManager`. It manages:

- `Books: map[int]*Book`
- `Members: map[int]*Member`
- Auto-incrementing `nextBookID` and `nextMemberID`

### 🧑‍💻 `LibraryController` Struct

Located in `controllers/library_controller.go`, it:

- Displays CLI menus
- Reads user input
- Calls service methods (e.g., `BorrowBook`, `AddBook`)
- Handles user-friendly messaging and error output

---

## 🧪 Testing

A comprehensive unit test suite is included for the `Library` service in:

```
services/library_service_test.go
```

### Run tests:

```bash
go test ./services/
```

### Verbose output:

```bash
go test -v ./services/
```

---

## 🚀 How to Run

### 🛠️ Prerequisites:

- Go 1.18+ installed

### ▶️ Steps:

```bash
cd library_management/
go mod tidy        # If needed
go run main.go
```

---

## 💻 Usage Guide

When the application runs, it shows this menu:

```
Library Management System
1. Add Book
2. Remove Book
3. Borrow Book
4. Return Book
5. List Available Books
6. List Borrowed Books
7. Add Member
8. Exit
```

### Example Interaction:

```text
Choose an option: 7
Enter Member Name: Alice
Member added successfully with ID 1

Choose an option: 1
Enter Book Title: Go Programming
Enter Book Author: John Doe
Book added successfully with ID 1
```

---

## 💡 Features

### ✅ Member Management:

- Add Members
- Remove Members (optional)
- List Borrowed Books per Member

### ✅ Book Management:

- Add Books
- Remove Books (only if not borrowed)

### ✅ Core Operations:

- Borrow Book (if available)
- Return Book
- List Available Books
- List Borrowed Books by Member

---

## ❗ Error Handling

- **Custom errors**:

  - `ErrBookNotFound`
  - `ErrMemberNotFound`
  - `ErrBookAlreadyBorrowed`
  - `ErrBookNotBorrowed`

- **Input validation**:

  - Ensures only valid IDs and options are accepted
  - Prevents crashing on invalid input (e.g., non-numeric values)

---

## 📝 Notes

- All data is stored **in-memory**, and **will reset** when the app exits
- Uses **Go best practices** like interfaces, dependency injection, and error wrapping

---

## 📄 License

This project is open source and intended for **educational purposes**. Feel free to modify and expand it as you learn Go and software architecture.
