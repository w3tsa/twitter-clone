package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/w3tsa/twitter-clone/database"
)

func main() {
    database.ConnectDb()
    app := fiber.New()

    setupRoutes(app)

    app.Listen(":8080")
}