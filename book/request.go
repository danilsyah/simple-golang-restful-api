package book

import "encoding/json"

type BookRequest struct {
	Title       string      `json:"title" binding:"required"`
	Price       json.Number `json:"price" binding:"required,number"`
	Description string      `json:"description" binding:"required"`
	Rating      json.Number `json:"rating" binding:"required,number"`
	Author      string      `json:"author" binding:"required"`

	// Email string      `json:"email" binding:"required,email"`
	// SubTitle string `json:"sub_title"`
}
