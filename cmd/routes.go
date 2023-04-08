package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/w3tsa/twitter-clone/handlers"
)


func setupRoutes(app *fiber.App) {
	app.Get("/api/v1/users", handlers.GetUser)
	app.Post("/api/v1/users", handlers.CreateUser)
	app.Delete("/api/v1/users/:id", handlers.DeleteUser)
}