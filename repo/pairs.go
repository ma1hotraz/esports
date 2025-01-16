package repo

import "github.com/iloginow/esportsdifference/slips"

const PAIRS_RESULT_KEY = "PAIRS_RESULT"

func SavePairsResult(result slips.PairsResult) error {
	return saveJsonData(PAIRS_RESULT_KEY, result)
}

func GetPairsResult() (slips.PairsResult, error) {
	var target slips.PairsResult
	if err := getJsonData(PAIRS_RESULT_KEY, &target); err != nil {
		return target, err
	}
	return target, nil
}
