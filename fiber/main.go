// package main

// import (
// 	"log"

// 	"github.com/devesh/fiberrest/book"
// 	"github.com/gofiber/fiber/v2"
// )

// func helloWorld(c *fiber.Ctx) {
// 	c.SendString("hello world")
// }

// func setupRoutes(app *fiber.App) {
// 	app.Get("/", helloWorld)

// 	app.Get("/books", book.GetBooks)
// 	app.Get("/books/:id", book.GetBook)
// 	app.Post("/books", book.AddBook)
// 	app.Patch("/books/:id", book.UpdateBook)
// 	app.Delete("/books/:id", book.DeleteBook)
// }

// func main() {
// 	book.InitialMigration()
// 	app := fiber.New()
// 	setupRoutes(app)

// 	log.Fatal(app.Listen(":3000"))
// }
