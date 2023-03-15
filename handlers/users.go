package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/w3tsa/twitter-clone/database"
	"github.com/w3tsa/twitter-clone/models"
)

func GetUser(c *fiber.Ctx) error {
	users := []models.User{}
	
	database.DB.Db.Find(&users)

	return c.Status(200).JSON(fiber.Map{
		"users": users,
	})
}

func CreateUser(c *fiber.Ctx) error { 
	user := new(models.User) 

	if err := c.BodyParser(user); err!= nil { 
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	database.DB.Db.Create(&user)

	return c.Status(200).JSON(user)
}