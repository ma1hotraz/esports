package sleeper

import (
	"encoding/json"
	"fmt"
	"io"
	"net"
	"net/http"
	"strconv"
	"time"

	"github.com/iloginow/esportsdifference/dto"
	"github.com/iloginow/esportsdifference/esport"
	"github.com/iloginow/esportsdifference/settings"
	"github.com/sirupsen/logrus"
)

var playerMap map[string]string

func init() {
	playerMap = GetCsgoPlayerMap()
}

func GetData() (*Data, error) {
	endpoint := settings.GetSleeperApi()
	transport := &http.Transport{
		Dial: (&net.Dialer{
			Timeout:   100 * time.Second,
			KeepAlive: 100 * time.Second,
		}).Dial,
		TLSHandshakeTimeout:   100 * time.Second,
		ResponseHeaderTimeout: 100 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
	}
	client := http.Client{
		Transport: transport,
		Timeout:   100 * time.Second,
	}
	data := Data{}
	req, reqErr := http.NewRequest("GET", endpoint, nil)
	if reqErr != nil {
		return &data, reqErr
	}
	res, resErr := client.Do(req)
	if resErr != nil {
		return &data, resErr
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			logrus.Error(err)
		}
	}(res.Body)
	if res.StatusCode != 200 {
		return &data, fmt.Errorf("Response pri status: %s", res.Status)
	}
	if jsonErr := json.NewDecoder(res.Body).Decode(&data); jsonErr != nil {
		return &data, jsonErr
	}
	return &data, nil
}

func GetCsgoPlayerMap() map[string]string {
	resp, err := http.Get("https://api.sleeper.app/players/cs")
	if err != nil {
		logrus.Error("Error fetching player data:", err)
		return nil
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		logrus.Error("Error: received non-200 response code")
		return nil
	}

	var player PlayerData
	var playerMap = make(map[string]string)
	if err := json.NewDecoder(resp.Body).Decode(&player); err != nil {
		logrus.Error("Error decoding player data: ", err)
		return nil
	}
	for _, player := range player {
		playerMap[player.PlayerID] = player.Metadata.Username
	}
	return playerMap
}

func convertToFloat(s string) float64 {
	multiplier, err := strconv.ParseFloat(s, 64)
	if err != nil {
		logrus.Error("Error parsing PayoutMultiplier: ", err)
		return 0
	}
	return multiplier
}

func (d *Data) Filter() []dto.SleeperRelevantData {
	data := []dto.SleeperRelevantData{}
	gameInDay := getAllCsgoGameInDay()
	for _, line := range *d {
		stat := line.GetStatType()
		game := getGameTime(gameInDay, line.GameID)
		playerName := getPlayerName(line.SubjectID)
		if len(line.Options) == 0 || game.GameID == "" || playerName == "" {
			continue
		}
		overUnder := ""
		if len(line.Options) == 1 {
			overUnder = line.Options[0].PayoutMultiplier + " " + string(line.Options[0].Outcome)
		}
		if len(line.Options) == 2 {
			overUnder = line.Options[0].PayoutMultiplier + " " + string(line.Options[0].Outcome) + "; " + line.Options[1].PayoutMultiplier + " " + string(line.Options[1].Outcome)
		}
		data = append(data, dto.SleeperRelevantData{
			StatType:  stat,
			Sport:     esport.CSGO,
			Player:    playerName,
			Time:      time.Unix(game.StartTime/1000, 0),
			Value:     line.Options[0].OutcomeValue,
			OverUnder: overUnder,
			Multiplier: func() float64 {
				multiplier, err := strconv.ParseFloat(line.Options[0].PayoutMultiplier, 64)
				if err != nil {
					logrus.Error("Error parsing PayoutMultiplier: ", err)
					return 0
				}
				return multiplier
			}(),
		})
	}
	return data
}

func getGameTime(gameInDay GameInDay, gameId string) GameData {
	for _, game := range gameInDay {
		if game.GameID == gameId {
			return game
		}
	}
	return GameData{}
}

func getPlayerName(playerId string) string {
	if playerMap == nil {
		return ""
	}
	return playerMap[playerId]
}

func getAllCsgoGameInDay() GameInDay {
	// get all csgo games in a day
	// get today in yyyy-mm-dd format
	today := time.Now().Format("2006-01-02")
	tommorow := time.Now().AddDate(0, 0, 1).Format("2006-01-02")
	nextday := time.Now().AddDate(0, 0, 2).Format("2006-01-02")
	nextday1 := time.Now().AddDate(0, 0, 3).Format("2006-01-02")
	nextday2 := time.Now().AddDate(0, 0, 4).Format("2006-01-02")
	dayList := []string{today, tommorow, nextday, nextday1, nextday2}
	var res GameInDay

	for _, day := range dayList {
		resp, err := http.Get("https://api.sleeper.app/scores/cs/date/" + day)
		if err != nil {
			logrus.Error("Error fetching player data:", err)
			return nil
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			logrus.Error("Error: received non-200 response code")
			return nil
		}

		var gameInDay GameInDay
		if err := json.NewDecoder(resp.Body).Decode(&gameInDay); err != nil {
			logrus.Error("Error decoding player data: ", err)
			return nil
		}
		for _, g := range gameInDay {
			res = append(res, g)
		}
	}

	return res
}
