package repo

import (
	"fmt"
	"time"
)

type User struct {
	PlayerId string `json:"playerid"`
	TeamId   string `json:"teamid"`
}

func StoreUserInRedis(user User) error {
	key := fmt.Sprintf("user:%s", user.PlayerId)
	err := client.HSet(ctx, key, map[string]interface{}{
		"playerid": user.PlayerId,
		"teamid":   user.TeamId,
	}).Err()
	if err != nil {
		return err
	}
	client.Expire(ctx, key, 2*24*time.Hour)
	return nil
}

func GetUserFromRedis(playerID string) (*User, error) {
	key := fmt.Sprintf("user:%s", playerID)
	result, err := client.HGetAll(ctx, key).Result()
	if err != nil {
		return nil, err
	}
	if len(result) == 0 {
		return nil, nil
	}
	return &User{
		PlayerId: result["playerid"],
		TeamId:   result["teamid"],
	}, nil
}
