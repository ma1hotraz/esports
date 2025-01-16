package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Id           uint       `json:"id"`
	Name         string     `json:"name"`
	Email        string     `json:"email" gorm:"unique"`
	Password     []byte     `json:"-"`
	InviteCodeID uint       `json:"invite_code_id"`
	InviteCode   InviteCode `gorm:"foreignKey:InviteCodeID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	IsAdmin      bool       `json:"is_admin"`
}

type InviteCode struct {
	gorm.Model
	Id             uint      `json:"id"`
	Code           string    `json:"invite_code" gorm:"unique"`
	Used           bool      `json:"used"`
	ExpirationTime time.Time `json:"expiration_time"`
	IsInfinite     bool      `json:"is_infinite"`
	Users          []User    `json:"users"`
}

type ForgotPasswordCode struct {
	gorm.Model
	Id             uint      `json:"id"`
	Code           string    `json:"invite_code" gorm:"unique"`
	Used           bool      `json:"used"`
	ExpirationTime time.Time `json:"expiration_time"`
	UserEmail      string    `json:"email"`
}

type Stat struct {
	gorm.Model
	Id        uint      `gorm:"primaryKey"`
	UserId    uint      `json:"user_id"`
	User      User      `json:"-" gorm:"foreignKey:UserId"`
	Player    string    `json:"player"`
	StatType  string    `json:"stat_type"`
	Timestamp time.Time `json:"timestamp"`
	Team      string    `json:"team"`
	Opponent  string    `json:"opponent"`
	Sport     string    `json:"sport"`
}

type Notification struct {
	gorm.Model
	Id      uint   `json:"notification_id" gorm:"primaryKey"`
	Message string `json:"message"`
}

type NotificationDismiss struct {
	gorm.Model
	Id             uint         `gorm:"primaryKey"`
	UserId         uint         `json:"user_id"`
	NotificationId uint         `json:"notification_id"`
	User           User         `json:"-" gorm:"foreignKey:UserId"`
	Notification   Notification `json:"-" gorm:"foreignKey:NotificationId"`
}
