package repo

import "github.com/iloginow/esportsdifference/slips"

const SLIPS_RESULT_KEY = "SLIPS_RESULT"

func SaveSlipsResult(result slips.Result) error {
	return saveJsonData(SLIPS_RESULT_KEY, result)
}

func GetSlipsResult() (slips.Result, error) {
	var target slips.Result
	if err := getJsonData(SLIPS_RESULT_KEY, &target); err != nil {
		return target, err
	}
	return target, nil
}
