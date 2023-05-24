package handlers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/w3tsa/twitter-clone/database"
	"github.com/w3tsa/twitter-clone/models"
)

type CreateTweetRequest struct {
	Content string
}

// Create a new post for a user 
func CreatePost(c *fiber.Ctx) error {
	// parse request body
	var req CreateTweetRequest
	if err := c.BodyParser(&req); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request payload"})
    }


	// Get userID from route parameters
	userID, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid user ID"})
	}

	// Retrieve user from the database
	user := models.User{}
	result := database.DB.Db.First(&user, userID)
	if result.Error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "User not found"})
	}

	// Create new post
	post := models.Post{
		Content: req.Content,
		UserID: user.ID,
		User: user,
	}

	// Save the post to the database
	if err := database.DB.Db.Create(&post).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create tweet"})
	}
	

	return c.Status(201).JSON(post)

}

// Get posts from the database with a specific id
func GetPost(c *fiber.Ctx) error {
	postId, err := strconv.Atoi(c.Params("postId"))

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	post := models.Post{}
	result := database.DB.Db.First(&post, postId)

	if result.Error != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Post not found"})
	}

	return c.Status(200).JSON(post)
}

// Update post of an existing post
func UpdatePost(c *fiber.Ctx) error {
	postId, err := strconv.Atoi(c.Params("postId"))

	if err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid post ID"})
    }

	post := models.Post{}
	result := database.DB.Db.First(&post, postId)

	if result.Error != nil {
        return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Post not found"})
    }

	if err := c.BodyParser(&post); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request payload"})
	}

	result = database.DB.Db.Save(&post)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to update post"})
	}

	return c.Status(200).JSON(post)
}

// Delete a post for a specific user id
func DeletePost(c *fiber.Ctx) error {
	userID, err := strconv.Atoi(c.Params("id"))

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid user ID"})
	}

	postId, err := strconv.Atoi(c.Params("postId"))

	if err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid post ID"})
    }

	user := models.User{}
	result := database.DB.Db.First(&user, userID)
	if result.Error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "User not found"})
	}

	post := models.Post{}
	result = database.DB.Db.First(&post, postId)

	if result.Error != nil {
        return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Post not found"})
    }


	database.DB.Db.Delete(&post)

	return c.Status(fiber.StatusNoContent).JSON(post)
 	
}