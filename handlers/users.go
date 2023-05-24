package handlers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/w3tsa/twitter-clone/database"
	"github.com/w3tsa/twitter-clone/models"
)

// Returns all the users in the database
func GetUser(c *fiber.Ctx) error {
	users := []models.User{}
	
	database.DB.Db.Find(&users)

	return c.Status(200).JSON(fiber.Map{
		"users": users,
	})
}

// Returns a specific user in the database by ID
func GetUserByID(c *fiber.Ctx) error {
	userID, err := strconv.Atoi(c.Params("id"))

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid user ID"})
	}

	user := models.User{}
	result := database.DB.Db.First(&user, userID)
	if result.Error != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "User not found"})
	}

	return c.Status(200).JSON(user)
}

// Creates a new user in the database
func CreateUser(c *fiber.Ctx) error { 
	user := new(models.User)

	if err := c.BodyParser(user); err!= nil { 
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	database.DB.Db.Create(&user)


	return c.Status(201).JSON(user)
}

// Updates an existing user
func UpdateUser(c *fiber.Ctx) error {
	userID, err := strconv.Atoi(c.Params("id"))

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid user ID"})
	}

	user := models.User{}
	result := database.DB.Db.First(&user, userID)
	if result.Error != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "User not found"})
	}

	if err := c.BodyParser(&user); err!= nil { 
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	result = database.DB.Db.Save(&user)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to update user"})
	}

	return c.Status(200).JSON(user)
}

// Delete a user from the database
func DeleteUser(c *fiber.Ctx) error {
	userID, err := strconv.Atoi(c.Params("id"))

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid user ID"})
	}

	user := models.User{}
	result := database.DB.Db.First(&user, userID)
	if result.Error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "User not found"})
	}
	
	database.DB.Db.Delete(&user)

	return c.Status(fiber.StatusNoContent).JSON(user)
 	
}
