package notifier

import (
	"bytes"
	"encoding/json"
	"net/http"
	"os"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/iloginow/esportsdifference/dto"
	"github.com/iloginow/esportsdifference/esport"
	"github.com/sirupsen/logrus"
)

var (
	session   *discordgo.Session
	channelID []string
)

func Init(t string, c []string) {
	var err error
	channelID = c
	getDiscordBots()
	getDiscordBotsV2()
	if t != "" {
		session, err = discordgo.New("Bot " + t)
		if err != nil {
			logrus.Fatal(err)
		}
	} else {
		logrus.Warn("Discord token was not provided. Notifications are disabled.")
	}
	if len(c) == 0 {
		logrus.Warn("Discord channel ID was not provied. Notifications are disabled.")
	}
}

func AnnounceNewLines(u NewUnderdogData, p NewPrizepicksData, c NewLinesData, sleeper []dto.SleeperRelevantData) {
	u.announce(UNDERDOG)
	p.announce(PRIZEPICKS)
	c.announce(COMPARE)
	announceNewLinesForLeagueV2(len(sleeper), esport.CSGO, SLEEPER)
}

func sendDiscordMessage(webhookURL string, message string) error {
	payload := map[string]string{"content": message}
	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", webhookURL, bytes.NewBuffer(payloadBytes))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNoContent {
		logrus.Debugf("Message sent successfully!\n")
	} else {
		logrus.Errorf("Failed to send message. Status code %d\n", resp.StatusCode)
	}

	return nil
}

func SendDiscordMessageToAllBotV2(message string) {
	for botName, webhookURL := range getDiscordBotsV2() {
		logrus.Debugf("Sending message by %s with content %s\n", botName, message)
		err := sendDiscordMessage(webhookURL, message)
		if err != nil {
			logrus.Errorf("Error: %v\n", err)
		}
	}
}

func getDiscordBots() map[string]string {
	prefix := "DISCORD_BOT_V1_"
	discordBotEnvVars := make(map[string]string)

	// Iterate over all environment variables
	for _, env := range os.Environ() {
		// Split each environment variable by '=' to get key and value
		parts := strings.SplitN(env, "=", 2)
		key := parts[0]
		value := parts[1]
		// Check if the key starts with the desired prefix
		if strings.HasPrefix(key, prefix) {
			logrus.Debugf("Bot: %s with url %s\n", key, value)
			// Store the environment variable with the prefix removed
			discordBotEnvVars[key] = value
		}
	}
	return discordBotEnvVars
}

func getDiscordBotsV2() map[string]string {
	prefix := "DISCORD_BOT_V2_"
	discordBotEnvVars := make(map[string]string)

	// Iterate over all environment variables
	for _, env := range os.Environ() {
		// Split each environment variable by '=' to get key and value
		parts := strings.SplitN(env, "=", 2)
		key := parts[0]
		value := parts[1]
		// Check if the key starts with the desired prefix
		if strings.HasPrefix(key, prefix) {
			logrus.Debugf("Bot: %s with url %s\n", key, value)
			// Store the environment variable with the prefix removed
			discordBotEnvVars[key] = value
		}
	}
	return discordBotEnvVars
}
