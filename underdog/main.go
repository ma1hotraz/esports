package underdog

import (
	"encoding/json"
	"fmt"
	"io"
	"net"
	"net/http"
	"os"
	"time"

	"github.com/iloginow/esportsdifference/dto"
	"github.com/iloginow/esportsdifference/settings"
	"github.com/sirupsen/logrus"
)

func GetData() (*Data, error) {
	endpoint := settings.GetUnderDogApi()

	transport := &http.Transport{
		Dial: (&net.Dialer{
			Timeout:   100 * time.Second,
			KeepAlive: 100 * time.Second,
		}).Dial,
		TLSHandshakeTimeout:   100 * time.Second,
		ResponseHeaderTimeout: 100 * time.Second,
		ExpectContinueTimeout: 100 * time.Second,
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
	req.Header.Add(
		"User-Agent",
		"Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/121.0.0.0 Safari/537.36",
	)
	res, resErr := client.Do(req)
	if resErr != nil {
		logrus.Infof("Response: %v\n", resErr)
		return &data, resErr
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			logrus.Error(err)
		}
	}(res.Body)
	if res.StatusCode != 200 {
		return &data, fmt.Errorf("response underdog status: %s", res.Status)
	}
	if jsonErr := json.NewDecoder(res.Body).Decode(&data); jsonErr != nil {
		// logrus.Infof("1 Response: %v\n", data)
		logrus.Infof("1 Eror: %v\n", jsonErr)
		return &data, jsonErr
	}
	return &data, nil
}

func GetDataFromJsonFile(filePath string) (*Data, error) {

	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Read the file contents
	bytes, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	// Unmarshal the JSON data into a Data struct
	var data Data
	if err := json.Unmarshal(bytes, &data); err != nil {
		return nil, err
	}

	return &data, nil
}

func IsRecordInList(record dto.UnderdogRelevantData, recordList []dto.UnderdogRelevantData) bool {
	for _, r := range recordList {
		if IsRecordsEqual(r, record) {
			return true
		}
	}
	return false
}

func IsRecordsEqual(r1, r2 dto.UnderdogRelevantData) bool {
	return r1.Player == r2.Player && r1.StatType == r2.StatType && r1.Team == r2.Team && r1.Opponent == r2.Opponent
}
