package scheduler

import (
	"context"
	"time"

	"codnect.io/chrono"
	"github.com/iloginow/esportsdifference/database"
	"github.com/iloginow/esportsdifference/models"
	"github.com/iloginow/esportsdifference/update"
	"github.com/sirupsen/logrus"
)

func scheduleUpdates() {
	s := chrono.NewDefaultTaskScheduler()
	_, err := s.ScheduleAtFixedRate(func(ctx context.Context) {
		update.SyncNewDifferencesData()
	}, 1*time.Minute)

	if err != nil {
		logrus.Fatal(err)
	}
}

func cleanStaleInformed() {
	s := chrono.NewDefaultTaskScheduler()
	_, err := s.ScheduleAtFixedRate(func(ctx context.Context) {
		update.CleanStaleInformed()

	}, time.Hour*12)

	if err != nil {
		logrus.Fatal(err)
	}
}

func cleanFWCode() {
	s := chrono.NewDefaultTaskScheduler()
	_, err := s.ScheduleAtFixedRate(func(ctx context.Context) {
		if (database.DB) == nil {
			return
		}
		if res := database.DB.Where("used = ? OR expiration_time < ?", true, time.Now()).Delete(&models.ForgotPasswordCode{}); res.Error != nil {
			logrus.Errorf("Error clean stale forgot password code")
		} else {
			logrus.Infof("Cleaned %d stale forgot passwords", res.RowsAffected)
		}

		if res2 := database.DB.Where("created_at < ?", time.Now().Add(time.Duration(-36)*time.Hour)).Delete(&models.Stat{}); res2.Error != nil {
			logrus.Errorf("Error clean stale used lines")
		} else {
			logrus.Infof("Cleaned %d stale used lines", res2.RowsAffected)
		}
	}, time.Hour*6)

	if err != nil {
		logrus.Fatal(err)
	}

}
