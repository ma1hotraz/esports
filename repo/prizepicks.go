package repo

import "github.com/iloginow/esportsdifference/prizepicks"

const PRIZEPICKS_RELEVANT_KEY = "PRIZEPICKS_RELEVANT"

func SavePrizepicksRelevant(data []prizepicks.RelevantData) error {
	return saveJsonData(PRIZEPICKS_RELEVANT_KEY, data)
}

func GetPrizepicksRelevant() ([]prizepicks.RelevantData, error) {
	var target []prizepicks.RelevantData
	if err := getJsonData(PRIZEPICKS_RELEVANT_KEY, &target); err != nil {
		return target, err
	}
	return target, nil
}
