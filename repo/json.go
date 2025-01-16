package repo

import (
	"encoding/json"

	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
)

func saveJsonData(key string, data interface{}) error {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}
	err = client.Set(ctx, key, jsonData, 0).Err()
	if err != nil {
		return err
	}
	return nil
}

func getJsonData(key string, target interface{}) error {
	jsonData, err := client.Get(ctx, key).Result()
	if err == redis.Nil {
		return nil
	} else if err != nil {
		return nil
	}
	err = json.Unmarshal([]byte(jsonData), target)
	if err != nil {
		return nil
	}
	return nil
}

func removeJsonData(key string) error {
	// Remove the value associated with the key.
	err := client.Del(ctx, key).Err()
	if err != nil {
		logrus.Fatalf("Error removing value for key %s: %v", key, err)
		return err
	}
	logrus.Infof("Value for key %s removed successfully\n", key)
	return nil
}
