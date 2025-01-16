package repo

import (
	"github.com/iloginow/esportsdifference/dto"
)

const Sleeper_RELEVANT_KEY = "Sleeper_RELEVANT_KEY"

func SaveSleeperRelevant(data []dto.SleeperRelevantData) error {
	return saveJsonData(Sleeper_RELEVANT_KEY, data)
}

func GetSleeperRelevant() ([]dto.SleeperRelevantData, error) {
	var target []dto.SleeperRelevantData
	if err := getJsonData(Sleeper_RELEVANT_KEY, &target); err != nil {
		return target, err
	}
	return target, nil
}
