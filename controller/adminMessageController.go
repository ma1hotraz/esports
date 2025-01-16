package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/iloginow/esportsdifference/database"
	"github.com/iloginow/esportsdifference/models"
)

func CreateNewNotification(c *fiber.Ctx) error {
	var notif models.Notification
	if err := c.BodyParser(&notif); err != nil {
		return err
	}

	// Insert into the database
	result := database.DB.Create(&notif)
	if result.Error != nil {
		return result.Error
	}

	return c.JSON(fiber.Map{
		"message": "Notification data inserted successfully",
	})
}

func GetNotifications(c *fiber.Ctx) error {
	var notifications []models.Notification

	database.DB.Table("notifications").Order("created_at desc").Find(&notifications)

	return c.JSON(fiber.Map{
		"notifications": notifications,
	})

}

func DeleteNotification(c *fiber.Ctx) error {
	notifId := c.Params("id")
	result := database.DB.Where("id = ?", notifId).Delete(&models.Notification{})
	if result.Error != nil {
		return result.Error
	}
	return c.JSON(fiber.Map{
		"message": "Notification with id" + notifId + "data inserted successfully",
	})
}

func DismissNotification(c *fiber.Ctx) error {
	var requestData struct {
		UserId         uint `json:"user_id"`
		NotificationId uint `json:"notification_id"`
	}
	if err := c.BodyParser(&requestData); err != nil {
		return err
	}

	// Check if the provided user ID is valid
	if requestData.UserId == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid user Id"})
	}

	if requestData.NotificationId == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid stat Id"})
	}

	notifDismiss := models.NotificationDismiss{
		UserId:         requestData.UserId,
		NotificationId: requestData.NotificationId,
	}
	if err := database.DB.Where("user_id = ? AND notification_id = ?", requestData.UserId, requestData.NotificationId).First(&notifDismiss).Error; err == nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Notification already dismissed"})
	}

	result := database.DB.Create(&notifDismiss)
	if result.Error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Notification dismiss data inserted successfully",
			"error":   result.Error.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"message": "Notification dismiss data inserted successfully",
	})
}

func GetMyNotifications(c *fiber.Ctx) error {
	userId := c.Params("userId")

	if userId == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid userId"})
	}

	var notifications []models.Notification

	database.DB.Table("notifications").Joins("LEFT JOIN notification_dismisses ON notifications.id = notification_dismisses.notification_id AND notification_dismisses.user_id = ?", userId).
		Where("notification_dismisses.notification_id IS NULL").Order("created_at desc").
		Find(&notifications)
	return c.JSON(fiber.Map{
		"notifications": notifications,
	})
}
