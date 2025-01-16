package underdog

import (
	"regexp"
	"strings"
	"time"

	"github.com/iloginow/esportsdifference/repo"
	"github.com/sirupsen/logrus"
)

type Game struct {
	AwayTeamId    string    `json:"away_team_id"`
	AwayTeamScore int       `json:"away_team_score"`
	HomeTeamId    string    `json:"home_team_id"`
	HomeTeamScore int       `json:"home_team_score"`
	Id            int       `json:"id"`
	MatchProgress string    `json:"match_progress"`
	Period        int       `json:"period"`
	ScheduledAt   time.Time `json:"scheduled_at"`
	SeasonType    string    `json:"season_type"`
	SportId       string    `json:"sport_id"`
	Status        string    `json:"status"`
	Title         string    `json:"title"`
	Type          string    `json:"type"`
	Year          int       `json:"year"`
}

type SoloGame struct {
	CompetitionId          string    `json:"competition_id"`
	Id                     int       `json:"id"`
	MatchProgress          string    `json:"match_progress"`
	Period                 int       `json:"period"`
	ScheduledAt            time.Time `json:"scheduled_at"`
	Score                  string    `json:"score"`
	SportId                string    `json:"sport_id"`
	SportTournamentRoundId int       `json:"sport_tournament_round_id"`
	Status                 string    `json:"status"`
	Title                  string    `json:"title"`
	Type                   string    `json:"type"`
}

func (game Game) GetTeamNamesByPlayer(player Player) (string, string) {
	// TODO: ADAPT
	parts := strings.Split(game.Title, ": ")
	var teams [2]string
	if len(parts) == 1 {
		teams = ExtractStrings(parts[0])
	} else {
		teams = ExtractStrings(parts[1])
	}
	// fmt.Println("teams", teams)
	if len(teams) != 2 {
		return "", ""
	}

	playerTeamId := player.TeamId

	if playerTeamId != "" {
		repo.StoreUserInRedis(repo.User{PlayerId: player.Id, TeamId: playerTeamId})
	} else {
		user, err := repo.GetUserFromRedis(player.Id)
		if err != nil {
			return "", ""
		}
		if user != nil {
			logrus.Debugf("Auto fix team id for player %v", player)
			playerTeamId = user.TeamId
		}
	}

	if playerTeamId == game.AwayTeamId {
		return teams[0], teams[1]
	}
	if playerTeamId == game.HomeTeamId {
		return teams[1], teams[0]
	}
	return "", ""
}

func ExtractStrings(input string) [2]string {
	// Define regular expression patterns for possible separators
	separators := []string{"v\\.s\\.", "vs\\.", "Vs\\.", "vs", "Vs", "VS", "VS\\.", "V\\.S\\."}

	// Create a regular expression pattern by joining the separators
	separatorPattern := strings.Join(separators, "|")

	// Create a regular expression
	re := regexp.MustCompile(separatorPattern)

	// Split the string using the regular expression pattern
	parts := re.Split(input, -1)

	// Trim spaces from each part
	for i := range parts {
		parts[i] = strings.TrimSpace(parts[i])
	}

	// Ensure only first two parts are considered
	if len(parts) >= 2 {
		parts = parts[:2]
	} else {
		return [2]string{"", ""}
	}

	// Return an array of length 2 containing the extracted substrings
	return [2]string{parts[0], parts[1]}
}
