package repo

import "github.com/iloginow/esportsdifference/compare"

const COMPARE_RESULT_KEY = "COMPARE_RESULT"
const SLEEPER_RESULT_KEY = "SLEEPER_RESULT_KEY"

func SaveCompareResult(result compare.Result) error {
	return saveJsonData(COMPARE_RESULT_KEY, result)
}

func SaveSleeperResult(result compare.Result) error {
	return saveJsonData(SLEEPER_RESULT_KEY, result)
}

func GetCompareResult() (compare.Result, error) {
	var target compare.Result
	if err := getJsonData(COMPARE_RESULT_KEY, &target); err != nil {
		return target, err
	}
	return target, nil
}

func getSleeperResult() (compare.Result, error) {
	var target compare.Result
	if err := getJsonData(SLEEPER_RESULT_KEY, &target); err != nil {
		return target, err
	}
	return target, nil
}
