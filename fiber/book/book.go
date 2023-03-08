package book

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Books struct {
	gorm.Model
	Name   string `JSON:"name"`
	Author string `JSON:"author"`
	Place  string `JSON:"place"`
	Rating int    `JSON:"rating"`
}

var DB *gorm.DB
var err error

const url = "postgres://fiberuser:fiber@123@localhost:5432/fiberapi?sslmode=disable"

func InitialMigration() {
	DB, err = gorm.Open(postgres.Open(url), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
		fmt.Println("Data connection failed")
	}
	DB.AutoMigrate(&Books{})
	fmt.Println("connected to database")
}

func GetBooks(c *fiber.Ctx) {
	// c.SendString("all book")\
	var book []Books
	DB.Find(&book)
	c.JSON(book)

}

func AddBook(c *fiber.Ctx) {
	// c.SendString("create book")
	var book Books
	book.Name = "spy"
	book.Author = "dee"
	book.Place = "India"
	book.Rating = 5
	DB.Create(&book)
	c.JSON(book)
}

func GetBook(c *fiber.Ctx) {
	// c.SendString("single book")
	var book Books
	id := c.Params("id")
	DB.Find(&book, id)
	c.JSON(book)
}

func UpdateBook(c *fiber.Ctx) {
	// c.SendString("update book")
	var book Books
	id := c.Params("id")
	DB.First(&book, id)
	if book.Author == " " {
		c.Status(500).SendString("Book  is not available")
	}
	if err := c.BodyParser(&book); err != nil {
		c.Status(500).SendString(err.Error())
	}
	DB.Save(&book)
	c.JSON(&book)
	c.SendString("book updated")

}

func DeleteBook(c *fiber.Ctx) {
	// c.SendString("delete book")
	var book Books
	id := c.Params("id")
	DB.First(&book, id)
	if book.Author == " " {
		c.Status(500).SendString("Book  is not available")
	}
	DB.Delete(&book)
	c.JSON("Book has been deleted")

}
