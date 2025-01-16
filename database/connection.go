package database

import (
	"fmt"

	"github.com/iloginow/esportsdifference/models"
	"github.com/iloginow/esportsdifference/utils"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {

	username := utils.GetEnvOrDefault("DB_USERNAME", "root")
	password := utils.GetEnvOrDefault("DB_PASSWORD", "root")
	host := utils.GetEnvOrDefault("DB_HOST", "localhost")
	port := utils.GetEnvOrDefault("DB_PORT", "5434")
	dbname := utils.GetEnvOrDefault("DB_NAME", "esportsdifference")

	// Construct DSN (Data Source Name) for database connection
	// dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", username, password, host, port, dbname)

	dsn := "host=%s user=%s password=%s dbname=%s port=%s sslmode=disable"
	dsn = fmt.Sprintf(dsn, host, username, password, dbname, port)

	connection, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("could not connect to the database")
	}

	DB = connection

	connection.AutoMigrate(&models.User{}, &models.InviteCode{}, &models.Stat{}, &models.ForgotPasswordCode{}, &models.Notification{}, &models.NotificationDismiss{})

	fmt.Println("Database connected")
}
