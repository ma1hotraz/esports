package repo

import "github.com/iloginow/esportsdifference/dto"

const UNDERDOGFANTAZY_RELEVANT_KEY = "UNDERDOGFANTAZY_RELEVANT"

func SaveUnderdogfantazyRelevant(data []dto.UnderdogRelevantData) error {
	return saveJsonData(UNDERDOGFANTAZY_RELEVANT_KEY, data)
}

func GetUnderdogfantazyRelevant() ([]dto.UnderdogRelevantData, error) {
	var target []dto.UnderdogRelevantData
	if err := getJsonData(UNDERDOGFANTAZY_RELEVANT_KEY, &target); err != nil {
		return target, err
	}
	return target, nil
}
