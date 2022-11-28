package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"

	"golang-restfulapi/book"
	"golang-restfulapi/handler"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:@tcp(127.0.0.1:3306)/db_pustaka?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Db Connection Error")
	}

	fmt.Println("========== Database connection succeed ============")

	db.AutoMigrate(&book.Book{})

	// ============== TEST ORM DB GORM ==========================

	// CRUD

	// ==================
	// create record
	// ==================
	// book := book.Book{}
	// book.Title = "Atomic habits"
	// book.Price = 75000
	// book.Description = "Buku self development tentang membangun kebiasaan baik dan menghilangkan kebiasaan buruk"
	// book.Author = "teu apal"
	// book.Rating = 4

	// err = db.Create(&book).Error
	// if err != nil {
	// 	fmt.Println("Error creating book record")
	// }

	// ===========================
	// retrieving single first object
	// ===========================
	// var book book.Book
	// err = db.Debug().First(&book, 2).Error
	// if err != nil {
	// 	fmt.Println("Error finding first book record")
	// }
	// fmt.Println("Title : ", book.Title)
	// fmt.Println("Author : ", book.Author)

	// =============================
	// retrieving single last object
	// =============================
	// var book book.Book
	// err = db.Debug().Last(&book).Error
	// if err != nil {
	// 	fmt.Println("Error finding last book record")
	// }
	// fmt.Println("Title : ", book.Title)
	// fmt.Println("Author : ", book.Author)

	// =============================
	// retrieving all object
	// =============================
	// var books []book.Book
	// err = db.Debug().Find(&books).Error
	// if err != nil {
	// 	fmt.Println("error finding all book record")
	// }
	// for _, b := range books {
	// 	fmt.Println("ID: ", b.ID)
	// 	fmt.Println("Title : ", b.Title)
	// 	fmt.Println("Description : ", b.Description)
	// 	fmt.Println("Author : ", b.Author)
	// 	fmt.Println("Price : ", b.Price)
	// 	fmt.Println("Rating : ", b.Rating)
	// 	fmt.Println("CreatedAt : ", b.CreatedAt)
	// 	fmt.Println("UpdatedAt : ", b.UpdatedAt)
	// }

	// ==================================
	// find by conditions
	// ==========================
	// var books []book.Book
	// // err = db.Debug().Where("author LIKE ?", "%danil%").Find(&books).Error
	// err = db.Debug().Where("rating = ?", 4).Find(&books).Error
	// if err != nil {
	// 	fmt.Println("error finding record")
	// }
	// for _, b := range books {
	// 	fmt.Println("ID: ", b.ID)
	// 	fmt.Println("Title : ", b.Title)
	// 	fmt.Println("Description : ", b.Description)
	// 	fmt.Println("Author : ", b.Author)
	// 	fmt.Println("Price : ", b.Price)
	// 	fmt.Println("Rating : ", b.Rating)
	// 	fmt.Println("CreatedAt : ", b.CreatedAt)
	// 	fmt.Println("UpdatedAt : ", b.UpdatedAt)
	// }

	// update record
	// var book book.Book
	// err = db.Debug().Where("id = ?", 1).First(&book).Error
	// if err != nil {
	// 	fmt.Println("Error finding book record")
	// }

	// book.Title = "Buku Doraemon"
	// err = db.Save(&book).Error
	// if err != nil {
	// 	fmt.Println("error updating book record")
	// }

	// delete
	// var book book.Book
	// err = db.Debug().Where("id = ?", 3).First(&book).Error
	// if err != nil {
	// 	fmt.Println("Error finding book record")
	// }

	// err = db.Delete(&book).Error
	// if err != nil {
	// 	fmt.Println("Error deleting book record")
	// }

	// ==========================================================

	// bookRepository := book.NewRepository(db)

	// book := book.Book{
	// 	Title:       "$100 Startup",
	// 	Description: "Good book",
	// 	Author:      "Udin",
	// 	Price:       250000,
	// 	Rating:      4,
	// }

	// bookRepository.Create(book)

	// books, err := bookRepository.FindAll()

	// if err != nil {
	// 	log.Fatal("error find books record")
	// }

	// for _, book := range books {
	// 	fmt.Println("Title : ", book.Title)
	// }

	// book, _ := bookRepository.FindByID(2)
	// fmt.Println("title : ", book.Title)

	// bookRequest := book.BookRequest{
	// 	Title: "buku motivasi hidup",
	// 	Price: "55000",
	// }

	// bookService.Create(bookRequest)

	bookRepository := book.NewRepository(db)
	// bookFileRepository := book.NewFileRepository()
	bookService := book.NewService(bookRepository)
	bookHandler := handler.NewBookHandler(bookService)

	router := gin.Default()

	v1 := router.Group("/v1")

	v1.GET("/books", bookHandler.ListAllBooksHandler)
	v1.GET("/books/:id", bookHandler.ListBookHandler)
	v1.POST("/books", bookHandler.PostBooksHandler)
	v1.PUT("/books/:id", bookHandler.UpdateBookHandler)
	v1.DELETE("/books/:id", bookHandler.DeleteBookHandler)

	// v1.GET("/", bookHandler.RootHandler)
	// v1.GET("/hello", bookHandler.HelloHandler)
	// v1.GET("/books/:id/:title", bookHandler.BooksHandler)
	// v1.GET("/query", bookHandler.QueryHandler)

	// v2 := router.Group("/v2")

	// v2.GET("/", bookHandler.RootHandler2)

	router.Run(":8888")

	// struktur project
	// main
	// handler
	// service
	// repository
	// db
	// mysql

}
