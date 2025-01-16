package prizepicks

import (
	"encoding/json"
	"fmt"
	"io"
	"net"
	"net/http"
	"os"
	"time"

	"github.com/iloginow/esportsdifference/settings"
	"github.com/sirupsen/logrus"
)

func GetData() (*Data, error) {
	endpoint := settings.GetPrizepicksApi()
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

func IsRecordInList(record RelevantData, recordList []RelevantData) bool {
	for _, r := range recordList {
		if IsRecordsEqual(r, record) {
			return true
		}
	}
	return false
}

func IsRecordsEqual(r1, r2 RelevantData) bool {
	return r1.Player == r2.Player && r1.StatType == r2.StatType && r1.Sport == r2.Sport
}
