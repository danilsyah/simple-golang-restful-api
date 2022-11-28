package book

import "time"

type Book struct {
	ID          int
	Title       string
	Description string
	Price       int
	Rating      int
	Author      string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
