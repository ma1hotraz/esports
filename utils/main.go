package utils

import (
	"fmt"
	"os"

	mailjet "github.com/mailjet/mailjet-apiv3-go"
	"github.com/sirupsen/logrus"
)

func GetEnvOrDefault(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func IsDevEnv() bool {
	return GetEnvOrDefault("ENV", "dev") == "dev"
}

var (
	SENDER        string
	mailjetClient *mailjet.Client
)

// Initialize email configuration
func InitEmailConfig() {
	SENDER = os.Getenv("MJ_SENDER_EMAIL")
	mjApiKeyPublic := os.Getenv("MJ_APIKEY_PUBLIC")
	mjApiKeyPrivate := os.Getenv("MJ_APIKEY_PRIVATE")
	if mjApiKeyPrivate == "" || mjApiKeyPublic == "" || SENDER == "" {
		logrus.Warnf("MJ APIs credentials are not provided. Sending email function is disabled")
		return
	}
	mailjetClient = mailjet.NewMailjetClient(mjApiKeyPublic, mjApiKeyPrivate)
}

func SendForgotPasswordEmail(sendTo, verificationCode string) error {
	if mailjetClient == nil {
		logrus.Warnf("Sending email function is not enabled")
		return nil
	}
	subject := "Forgot Password Assistance - Your Verification Code"
	body := fmt.Sprintf("Dear User,\n\nWe received a request to reset your password. Your verification code is: %s\n\nIf you didn't request this password reset, please disregard this email.\n\nBest regards,\nYour Service Team", verificationCode)
	htmlPart := fmt.Sprintf(`<html>
    <body>
        <p>Dear <strong>@%s</strong>,</p>
        <p>We received a request to reset your password. Your verification code is: <strong>%s</strong></p>
		<p>Please note that this code will expire in 30 minutes.</p>
        <p>If you didn't request this password reset, please disregard this email.</p>
        <p>Best regards,<br/>Esport Differences Team</p>
    </body>
    </html>`, sendTo, verificationCode)
	messagesInfo := []mailjet.InfoMessagesV31{
		mailjet.InfoMessagesV31{
			From: &mailjet.RecipientV31{
				Email: SENDER,
				Name:  "Esport Differences",
			},
			To: &mailjet.RecipientsV31{
				mailjet.RecipientV31{
					Email: sendTo,
					Name:  "Recipient",
				},
			},
			Subject:  subject,
			TextPart: body,
			HTMLPart: htmlPart,
		},
	}
	messages := mailjet.MessagesV31{Info: messagesInfo}
	res, err := mailjetClient.SendMailV31(&messages)
	if err != nil {
		return err
	}
	logrus.Infof("Sending email succeed with response data: %+v\n", res)

	return nil
}
