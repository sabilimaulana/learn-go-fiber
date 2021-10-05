package main

import (
	"fmt"
	"learn-fiber/book"
	"learn-fiber/database"
	"log"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	app := fiber.New()
	initDatabase()

	setupRoutes(app)

	log.Fatal(app.Listen(":8080"))
}

func setupRoutes(app *fiber.App) {
	app.Get("/books", book.GetBooks)
	app.Get("/books/:id", book.GetBook)
	app.Post("/books", book.NewBook)
	app.Delete("/books/:id", book.DeleteBook)
}

func initDatabase() {
	var err error
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	dsn := "root:@tcp(127.0.0.1:3306)/fiber-book?charset=utf8mb4&parseTime=True&loc=Local"
	database.DBConn, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println("Connection to database success")
	// database.DBConn.AutoMigrate(&book.Book{})
}
