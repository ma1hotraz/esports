package controllers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/iloginow/esportsdifference/database"
	"github.com/iloginow/esportsdifference/models"
)

func InsertNewUsed(c *fiber.Ctx) error {
	// Parse JSON body
	var stat models.Stat
	if err := c.BodyParser(&stat); err != nil {
		return err
	}

	// Insert into the database
	result := database.DB.Create(&stat)
	if result.Error != nil {
		return result.Error
	}

	return c.JSON(fiber.Map{
		"message": "Stat data inserted successfully",
		"stat":    stat,
	})
}

func RemoveAllUsed(c *fiber.Ctx) error {
	// Parse JSON body
	var requestData struct {
		UserID uint `json:"user_id"`
	}
	if err := c.BodyParser(&requestData); err != nil {
		return err
	}

	// Delete stats of the user from the database
	result := database.DB.Where("user_id = ?", requestData.UserID).Delete(&models.Stat{})
	if result.Error != nil {
		return result.Error
	}

	return c.JSON(fiber.Map{
		"message": fmt.Sprintf("Stats of user %d removed successfully", requestData.UserID),
	})
}

func GetAllUsedRows(c *fiber.Ctx) error {
	// Parse JSON body
	var requestData struct {
		UserID uint `json:"user_id"`
	}
	if err := c.BodyParser(&requestData); err != nil {
		return err
	}

	// Query stats of the user from the database
	var stats []models.Stat
	result := database.DB.Where("user_id = ?", requestData.UserID).Omit("User").Find(&stats)
	if result.Error != nil {
		return result.Error
	}

	return c.JSON(fiber.Map{
		"stats": stats,
	})
}

func RemoveOneUsed(c *fiber.Ctx) error {
	// Parse JSON body
	var requestData struct {
		UserID uint `json:"user_id"`
		StatID uint `json:"stat_id"`
	}
	if err := c.BodyParser(&requestData); err != nil {
		return err
	}

	// Check if the provided user ID is valid
	if requestData.UserID == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid user ID"})
	}

	// Check if the provided stat ID is valid
	if requestData.StatID == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid stat ID"})
	}

	// Remove the specific stat associated with the provided user ID and stat ID from the database
	result := database.DB.Where("user_id = ? AND id = ?", requestData.UserID, requestData.StatID).Delete(&models.Stat{})
	if result.Error != nil {
		return result.Error
	}

	return c.JSON(fiber.Map{
		"message": fmt.Sprintf("Stat with ID %d removed successfully for user %d", requestData.StatID, requestData.UserID),
	})
}
