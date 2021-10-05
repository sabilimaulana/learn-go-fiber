package book

import (
	"fmt"
	"learn-fiber/database"
	"time"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	ID        int
	Title     string `json:"title"`
	Author    string `json:"author"`
	Rating    int    `json:"rating"`
	UpdatedAt time.Time
	CreatedAt time.Time
}

type BookFormatter struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Author    string    `json:"author"`
	Rating    int       `json:"rating"`
	UpdatedAt time.Time `json:"updated_at"`
	CreatedAt time.Time `json:"created_at"`
}

func formatBook(book Book) BookFormatter {
	formatterBook := BookFormatter{
		ID:        book.ID,
		Title:     book.Title,
		Author:    book.Author,
		Rating:    book.Rating,
		CreatedAt: book.CreatedAt,
		UpdatedAt: book.UpdatedAt,
	}
	return formatterBook
}

func formatBooks(books []Book) []BookFormatter {
	var formatterBooks []BookFormatter

	for _, b := range books {
		formatterBook := BookFormatter{
			ID:        b.ID,
			Title:     b.Title,
			Author:    b.Author,
			Rating:    b.Rating,
			CreatedAt: b.CreatedAt,
			UpdatedAt: b.UpdatedAt,
		}
		formatterBooks = append(formatterBooks, formatterBook)
	}
	return formatterBooks
}

func GetBooks(c *fiber.Ctx) error {
	db := database.DBConn
	var books []Book
	err := db.Find(&books).Error
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	return c.JSON(formatBooks(books))
}

func GetBook(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DBConn
	var book Book

	err := db.Find(&book, id).Error
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	return c.JSON(formatBook(book))
}

func NewBook(c *fiber.Ctx) error {
	db := database.DBConn

	book := new(Book)
	if err := c.BodyParser(&book); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON("Add new book failed")
	}

	err := db.Create(&book).Error
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	return c.JSON(formatBook(*book))
}

func DeleteBook(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DBConn

	var book Book
	err := db.First(&book, id).Error
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	if book.ID == 0 {
		return c.Status(500).JSON("No book with that id")
	}

	err = db.Delete(&book).Error
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	return c.JSON("Book successfully deleted")
}
