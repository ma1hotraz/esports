package settings

import "github.com/iloginow/esportsdifference/utils"

func GetUnderDogApi() string {
	if utils.IsDevEnv() {
		return "https://api-underdogfantasy-com.translate.goog/beta/v5/over_under_lines?_x_tr_sl=en&_x_tr_tl=vi&_x_tr_hl=vi&_x_tr_pto=wapp"
	}
	return "https://api.underdogfantasy.com/beta/v5/over_under_lines"
}

func GetPrizepicksApi() string {
	if utils.IsDevEnv() {
		return "https://partner-api.prizepicks.com/projections"
	}
	// return "https://partner--api-prizepicks-com.translate.goog/projections?_x_tr_sl=en&_x_tr_tl=vi&_x_tr_hl=vi&_x_tr_pto=wapp"
	return "https://partner-api.prizepicks.com/projections"
}

func GetSleeperApi() string {
	return "https://api.sleeper.app/lines/available?sports[]=cs"
}
