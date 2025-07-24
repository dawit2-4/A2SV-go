package models

type Status string

const (
	Borrowed  Status = "borrowed"
	Available Status = "available"
)

type Book struct {
	ID     int
	Title  string
	Author string
	Status Status
}

func (s Status) IsAvailable() bool {
	return s == Available
}