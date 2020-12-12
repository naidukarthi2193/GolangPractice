package book

import (
	"encoding/json"
	"fmt"

	"github.com/TutorialEdge/go-fiber-rest-api-tutorial/database"
	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type Book struct {
	gorm.Model
	Title  string `json:"name"`
	Author string `json:"author"`
	Rating int    `json:"rating"`
}

func GetBooks(c *fiber.Ctx) {
	fmt.Println("Get Books")
	db := database.DBConn
	var books []Book
	db.Find(&books)
	c.JSON(books)
}

func GetBook(c *fiber.Ctx) {
	fmt.Println("Get Book")
	id := c.Params("id")
	db := database.DBConn
	var book Book
	db.Find(&book, id)
	c.JSON(book)
}

func NewBook(c *fiber.Ctx) {
	fmt.Println("New Book")
	db := database.DBConn
	var book Book
	fmt.Println(c.Body())
	err := json.Unmarshal([]byte(c.Body()), &book)
	if err != nil {
		panic(err)
	}
	book.Title = book.Title
	book.Author = book.Author
	book.Rating = book.Rating
	db.Create(&book)
	c.JSON(book)
}

func DeleteBook(c *fiber.Ctx) {
	id := c.Params("id")
	db := database.DBConn
	fmt.Println("Delete Book")

	var book Book
	db.First(&book, id)
	if book.Title == "" {
		c.Status(500).Send("No Book Found with ID")
		return
	}
	db.Delete(&book)
	c.Send("Book Successfully deleted")
}
