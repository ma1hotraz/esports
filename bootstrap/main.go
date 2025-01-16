package bootstrap

import (
	"github.com/iloginow/esportsdifference/database"
	"github.com/iloginow/esportsdifference/models"
	"github.com/iloginow/esportsdifference/utils"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

func RegisterInitAdminAccount() {
	// Check if the email is already used

	var adminEmail = utils.GetEnvOrDefault("BOOTSTRAP_ADMIN_EMAIL", "admin@admin.com")
	var adminPassword = utils.GetEnvOrDefault("BOOTSTRAP_ADMIN_PASSWORD", "test")

	if adminEmail == "" || adminPassword == "" {
		return
	}
	var existingUser models.User
	if err := database.DB.Preload("InviteCode").Where("email = ?", adminEmail).First(&existingUser).Error; err == nil {
		ivcode := existingUser.InviteCode
		ivcode.IsInfinite = true
		database.DB.Save(&ivcode)
		database.DB.Save(&existingUser)
		return
	}

	decryptedPassword, _ := bcrypt.GenerateFromPassword([]byte(adminPassword), 14)

	var inviteCode = models.InviteCode{
		Code:       "zQ8mKu5R7E3pWcDn",
		IsInfinite: true,
	}

	user := models.User{
		Name:       "Admin",
		Email:      adminEmail,
		Password:   decryptedPassword,
		IsAdmin:    true,
		InviteCode: inviteCode,
	}
	inviteCode.Used = true

	database.DB.Save(&inviteCode)

	database.DB.Create(&user)

	logrus.Info("Boostrap: Added new admin user")
}
