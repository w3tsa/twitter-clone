package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/w3tsa/twitter-clone/database"
)

func main() {
    database.ConnectDb()
    app := fiber.New()

    app.Get("/", func(c *fiber.Ctx) error {
        return c.SendString("This is a!")
    })

    app.Listen(":8080")
}