package controllers

import (
	"errors"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/iloginow/esportsdifference/database"
	"github.com/iloginow/esportsdifference/models"
	"github.com/iloginow/esportsdifference/utils"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

const SecretKey = "secret"

func encryptPassword(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), 14)
}

func Register(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	// Check if the email is already used
	var existingUser models.User
	if err := database.DB.Where("email = ?", data["email"]).First(&existingUser).Error; err == nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Email already in use"})
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}

	// Query the InviteCode table to retrieve the record for the provided invite code
	var inviteCode models.InviteCode
	if err := database.DB.Where("code = ?", data["invite_code"]).First(&inviteCode).Error; err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid or already used invite code"})
	}

	if inviteCode.Used {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invite code already used"})
	}

	password, _ := encryptPassword(data["password"])
	isAdmin, _ := strconv.ParseBool(data["is_admin"])

	user := models.User{
		Name:       data["name"],
		Email:      data["email"],
		Password:   password,
		InviteCode: inviteCode,
		IsAdmin:    isAdmin,
	}

	// Mark the invite code as used
	markUsedInviteCode(&inviteCode)

	database.DB.Save(&inviteCode)

	database.DB.Create(&user)

	return c.JSON(user)
}

func Login(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	var user models.User

	database.DB.Preload("InviteCode").Where("email = ?", data["email"]).First(&user)

	if user.Id == 0 {
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"message": "user not found",
		})
	}

	if user.InviteCode.IsInfinite == false && user.InviteCode.ExpirationTime.Before(time.Now()) {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "user was expired",
		})
	}

	if err := bcrypt.CompareHashAndPassword(user.Password, []byte(data["password"])); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "incorrect password",
		})
	}

	claims := jwt.MapClaims{
		"userId": user.Email,
		//set one hour expiration
		"exp": time.Now().Add(time.Second * 3).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(utils.GetEnvOrDefault("JWT_SECRET", "dumamay")))
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	cookie := fiber.Cookie{
		Name:     "jwtauth",
		Value:    t,
		Expires:  time.Now().Add(time.Hour * 24),
		SameSite: "none",
		HTTPOnly: false,
	}
	c.Cookie(&cookie)
	c.Cookie(&fiber.Cookie{
		Name:     "userId",
		Value:    strconv.FormatUint(uint64(user.Id), 10),
		Expires:  time.Now().Add(time.Hour * 24),
		SameSite: "none",
		HTTPOnly: false,
	})
	// Create refresh token
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId": user.Email,
		"exp":    time.Now().Add(time.Hour * 24).Unix(),
	})
	signedRefreshToken, err := refreshToken.SignedString([]byte(utils.GetEnvOrDefault("JWT_SECRET", "dumamay")))
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	// Prepare response with success message and isAdmin field
	response := fiber.Map{
		"message":         "success",
		"is_admin":        user.IsAdmin,
		"refresh_token":   string(signedRefreshToken),
		"email":           user.Email,
		"userId":          user.Id,
		"userData":        user,
		"expiration_time": user.InviteCode.ExpirationTime,
	}

	c.JSON(response)

	return c.SendStatus(fiber.StatusOK)
}

func CreateInviteCode(c *fiber.Ctx) error {
	type RequestInviteCode struct {
		InviteCode     string    `json:"invite_code"`
		ExpirationTime time.Time `json:"expiration_time"`
		IsInfinite     bool      `json:"is_infinite"`
	}

	var data RequestInviteCode
	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
	}

	inviteCode := data.InviteCode
	// Check if invite code already exists
	var existingInviteCode models.InviteCode
	if err := database.DB.Where("code = ?", inviteCode).First(&existingInviteCode).Error; err == nil {
		// Invite code already exists, return error
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invite code already exists"})
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		// Database error occurred
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Database error"})
	}

	// Create a new InviteCode record
	newInviteCode := models.InviteCode{
		Code:           inviteCode,
		ExpirationTime: data.ExpirationTime,
		IsInfinite:     data.IsInfinite,
	}

	// Save the new invite code
	if err := database.DB.Create(&newInviteCode).Error; err != nil {
		// Error occurred while saving the new invite code
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create invite code"})
	}

	// Return the newly created invite code
	return c.JSON(newInviteCode)
}

func User(c *fiber.Ctx) error {
	cookie := c.Cookies("userId")

	// token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
	// 	return []byte(SecretKey), nil
	// })

	// if err != nil {
	// 	c.Status(fiber.StatusUnauthorized)
	// 	return c.JSON(fiber.Map{
	// 		"message": "unauthenticated",
	// 	})
	// }

	// claims := token.Claims.(*jwt.StandardClaims)

	var user models.User
	if err := database.DB.Preload("InviteCode").Where("id = ?", cookie).First(&user).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "failed to get user",
		})
	}
	database.DB.Where("id = ?", cookie).First(&user)

	return c.JSON(user)
}

func GetAllUsers(c *fiber.Ctx) error {
	var users []models.User
	// Preload the InviteCode for each user
	if err := database.DB.Preload("InviteCode").Find(&users).Error; err != nil {
		return err
	}
	return c.JSON(users)
}

func GetUserById(c *fiber.Ctx) error {
	// Preload the InviteCode for each user
	userID := c.Params("id")
	var user models.User

	database.DB.Preload("InviteCode").Where("id = ?", userID).First(&user)

	return c.JSON(user)
}

func DeleteUser(c *fiber.Ctx) error {
	userID := c.Params("id")

	var user models.User
	if err := database.DB.Where("id = ?", userID).First(&user).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "user not found",
		})
	}

	if user.Email == "admin@admin.com" {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"message": "super user cannot be deleted",
		})
	}

	if err := database.DB.Delete(&user).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "failed to delete user",
		})
	}

	return c.JSON(fiber.Map{
		"message": "user deleted successfully",
	})
}

func EditUser(c *fiber.Ctx) error {
	userID := c.Params("id")
	type UpdatedUserRequest struct {
		Name           string    `json:"name"`
		Email          string    `json:"email"`
		Password       string    `json:"password"`
		ExpirationTime time.Time `json:"expiration_time"`
		IsAdmin        bool      `json:"is_admin"`
		IsInfinite     bool      `json:"is_infinite"`
	}

	var dto UpdatedUserRequest
	if err := c.BodyParser(&dto); err != nil {
		return err
	}

	var user models.User
	if err := database.DB.Preload("InviteCode").Where("id = ?", userID).First(&user).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "user not found",
		})
	}

	updatedUser := models.User{
		Name:    dto.Name,
		Email:   dto.Name,
		IsAdmin: dto.IsAdmin,
	}

	if dto.Password != "" {
		pw, _ := encryptPassword(dto.Password)
		updatedUser.Password = pw
	}

	inviteCode := user.InviteCode
	inviteCode.IsInfinite = dto.IsInfinite
	if !dto.IsInfinite {
		inviteCode.ExpirationTime = dto.ExpirationTime
	}

	database.DB.Save(&inviteCode)

	models.UpdateUserFields(&user, updatedUser)

	database.DB.Save(&user)

	return c.JSON(fiber.Map{
		"message": "user updated successfully",
		"user":    user,
	})
}

func GetAllInviteCodes(c *fiber.Ctx) error {
	var inviteCodes []models.InviteCode
	if err := database.DB.Preload("Users").Find(&inviteCodes).Error; err != nil {
		return err
	}
	return c.JSON(inviteCodes)
}

// func GetAllInviteCodesWithUsers(db *gorm.DB) ([]models.InviteCode, error) {
// 	var inviteCodes []models.InviteCode

// 	// Fetch all invite codes
// 	if err := db.Find(&inviteCodes).Error; err != nil {
// 		return nil, err
// 	}

// 	// Preload users for each invite code
// 	for i := range inviteCodes {
// 		if err := db.Model(&inviteCodes[i]).Association("Users").Find(&inviteCodes[i].); err != nil {
// 			return nil, err
// 		}
// 	}

// 	return inviteCodes, nil
// }

func Logout(c *fiber.Ctx) error {
	cookie := fiber.Cookie{
		Name:     "jwtauth",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HTTPOnly: true,
	}

	c.Cookie(&cookie)

	return c.JSON(fiber.Map{
		"message": "success",
	})
}

func DeleteInviteCode(c *fiber.Ctx) error {
	// Parse invite code ID from request parameters or request body
	inviteCodeID := c.Params("id")

	// Check if invite code exists
	var existingInviteCode models.InviteCode
	if err := database.DB.Where("id = ?", inviteCodeID).First(&existingInviteCode).Error; err != nil {
		// Invite code not found
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Invite code not found"})
	}

	// Delete the invite code from the database
	if err := database.DB.Delete(&existingInviteCode).Error; err != nil {
		// Error occurred while deleting invite code
		return err
	}

	// Invite code successfully deleted
	return c.JSON(fiber.Map{"message": "Invite code deleted successfully"})
}

func ForgotPassword(c *fiber.Ctx) error {
	var requestBody struct {
		Email string `json:"email"`
	}
	if err := c.BodyParser(&requestBody); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
	}

	// Check if the email exists in the users table
	var user models.User
	if err := database.DB.Where("email = ?", requestBody.Email).First(&user).Error; err != nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{"error": "User not found"})
	}

	var newForgotPasswordCode = models.ForgotPasswordCode{
		Code:           generateForgotPasswordCode(),
		UserEmail:      requestBody.Email,
		ExpirationTime: time.Now().Add(30 * time.Minute),
		Used:           false,
	}

	if err := database.DB.Save(&newForgotPasswordCode).Error; err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to save new forgot password code"})
	}
	utils.SendForgotPasswordEmail(requestBody.Email, newForgotPasswordCode.Code)
	return c.Status(http.StatusOK).JSON(fiber.Map{"message": fmt.Sprintf("Forgot password code sent to %s", requestBody.Email)})
}

func ChangeUserPw(c *fiber.Ctx) error {
	var requestBody struct {
		Email              string `json:"email"`
		ForgotPasswordCode string `json:"code"`
		NewPassword        string `json:"newPassword"`
	}
	if err := c.BodyParser(&requestBody); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
	}

	// Check if the email exists in the users table
	var user models.User
	if err := database.DB.Where("email = ?", requestBody.Email).First(&user).Error; err != nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{"error": "User not found"})
	}

	var fpCode models.ForgotPasswordCode

	if err := database.DB.Where("user_email = ? AND expiration_time > ? AND used = ? AND code = ?", requestBody.Email, time.Now(), false, requestBody.ForgotPasswordCode).First(&fpCode).Error; err != nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{"error": "Forgot password code was not found or expired"})
	}

	newPw, err := encryptPassword(requestBody.NewPassword)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to decrypt new password code"})
	}

	user.Password = newPw
	if err := database.DB.Save(&user).Error; err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to save new password"})
	}

	fpCode.Used = true
	if err := database.DB.Save(&fpCode).Error; err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to save new forgot password code"})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{"message": fmt.Sprintf("Password of %s changed", requestBody.Email)})
}

func ApplyNewInviteCode(c *fiber.Ctx) error {
	// Parse request body
	var requestBody struct {
		Email      string `json:"email"`
		InviteCode string `json:"invite_code"`
	}
	if err := c.BodyParser(&requestBody); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
	}

	// Check if the email exists in the users table
	var user models.User
	if err := database.DB.Where("email = ?", requestBody.Email).First(&user).Error; err != nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{"error": "User not found"})
	}

	// Check if the invite code exists in the invite codes table
	var inviteCode models.InviteCode
	if err := database.DB.Where("code = ?", requestBody.InviteCode).First(&inviteCode).Error; err != nil {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid invite code"})
	}

	// Check if the invite code is linked to another user
	if inviteCode.Used {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{"error": "Invite code is already linked to another user"})
	}
	// Check if the invite code is expired
	// if inviteCode.ExpirationTime.Before(time.Now()) {
	// 	return c.Status(http.StatusUnauthorized).JSON(fiber.Map{"error": "Invite code has expired"})
	// }

	markUsedInviteCode(&inviteCode)

	if err := database.DB.Save(&inviteCode).Error; err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to update invite code"})
	}
	user.InviteCode = inviteCode
	if err := database.DB.Save(&user).Error; err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to update expiration time"})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{"message": "Expiration time updated successfully"})
}

func markUsedInviteCode(inviteCode *models.InviteCode) {
	if inviteCode.Used {
		return
	}
	inviteCode.Used = true
	// if !inviteCode.IsInfinite {
	// 	inviteCode.ExpirationTime =
	// }
}

func generateForgotPasswordCode() string {
	// Generate a random 6-digit code
	code := fmt.Sprintf("%06d", rand.Intn(1000000))

	return code
}
