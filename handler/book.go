package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"

	"golang-restfulapi/book"
)

type bookHandler struct {
	bookService book.Service
}

func NewBookHandler(bookService book.Service) *bookHandler {
	return &bookHandler{bookService}
}

// func (h *bookHandler) RootHandler(c *gin.Context) {
// 	c.JSON(http.StatusOK, gin.H{
// 		"name": "Danil Syah",
// 		"bio":  "A Backend Developer",
// 	})
// }

// func (h *bookHandler) RootHandler2(c *gin.Context) {
// 	c.JSON(http.StatusOK, gin.H{
// 		"first_name": "danil",
// 		"last_name":  "syah",
// 		"age":        "27",
// 	})
// }

// func (h *bookHandler) HelloHandler(c *gin.Context) {
// 	c.JSON(http.StatusOK, gin.H{
// 		"intro": "hello world",
// 		"age":   2,
// 	})
// }

// func (h *bookHandler) BooksHandler(c *gin.Context) {
// 	id := c.Param("id")
// 	title := c.Param("title")

// 	c.JSON(http.StatusOK, gin.H{
// 		"id":    id,
// 		"title": title,
// 	})
// }

// func (h *bookHandler) QueryHandler(c *gin.Context) {
// 	title := c.Query("title")
// 	n1 := c.Query("n1")
// 	n2 := c.Query("n2")

// 	var result int

// 	var num1, _ = strconv.Atoi(n1)
// 	var num2, _ = strconv.Atoi(n2)

// 	result = num1 + num2

// 	var str = strconv.Itoa(result)

// 	c.JSON(http.StatusOK, gin.H{
// 		"title":  title,
// 		"n1":     n1,
// 		"n2":     n2,
// 		"result": str,
// 	})
// }

func (h *bookHandler) ListAllBooksHandler(c *gin.Context) {
	books, err := h.bookService.FindAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
	}

	var booksResponse []book.BookResponse

	for _, b := range books {
		bookResponse := convertToBookResponse(b)

		booksResponse = append(booksResponse, bookResponse)
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "berhasil menampilkan seluruh data",
		"result": booksResponse,
	})
}

func (h *bookHandler) ListBookHandler(c *gin.Context) {
	idString := c.Param("id")
	id, _ := strconv.Atoi(idString)

	book, err := h.bookService.FindByID(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	if book.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "data not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "berhasil menampilkan data",
		"result": convertToBookResponse(book),
	})

}

func (h *bookHandler) PostBooksHandler(c *gin.Context) {
	var bookRequest book.BookRequest

	err := c.ShouldBindJSON(&bookRequest)
	if err != nil {

		errorMessages := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error on field %s, condition: %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
		}

		c.JSON(http.StatusBadRequest, gin.H{
			"errors": errorMessages,
		})
		return

	}

	book, err := h.bookService.Create(bookRequest)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "berhasil menambahkan data",
		"result": convertToBookResponse(book),
	})
}

func (h *bookHandler) UpdateBookHandler(c *gin.Context) {
	var bookRequest book.BookRequest

	err := c.ShouldBindJSON(&bookRequest)
	if err != nil {
		errorMessages := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error on field %s, condition: %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
		}

		c.JSON(http.StatusBadRequest, gin.H{
			"errors": errorMessages,
		})
		return
	}

	idString := c.Param("id")
	id, _ := strconv.Atoi(idString)

	book, err := h.bookService.Update(id, bookRequest)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	if book.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Data not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "berhasil update data",
		"result": convertToBookResponse(book),
	})

}

func (h *bookHandler) DeleteBookHandler(c *gin.Context) {
	idString := c.Param("id")
	id, _ := strconv.Atoi(idString)

	_, err := h.bookService.Delete(id)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "data not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "berhasil menghapus data",
	})

}

func convertToBookResponse(b book.Book) book.BookResponse {
	return book.BookResponse{
		ID:          b.ID,
		Title:       b.Title,
		Description: b.Description,
		Price:       b.Price,
		Rating:      b.Rating,
		Author:      b.Author,
		CreatedAt:   b.CreatedAt,
		UpdatedAt:   b.UpdatedAt,
	}
}
