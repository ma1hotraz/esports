package logger

import (
	"io"
	"os"
	"path/filepath"
	"sync"
	"time"

	"github.com/iloginow/esportsdifference/utils"
	"github.com/sirupsen/logrus"
)

var (
	logFile *os.File
	once    sync.Once
)

func Init(logsDir string) {
	once.Do(func() {
		ensureLogsDir(logsDir)
		currentTime := time.Now()
		date := currentTime.Format("2006-01-02")
		logFileName := filepath.Join(logsDir, date+".log")
		f, err := os.OpenFile(logFileName, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
		if err != nil {
			logrus.Fatal(err)
		}
		logFile = f
		mw := io.MultiWriter(os.Stdout, logFile)
		logrus.SetFormatter(&logrus.TextFormatter{
			ForceColors:   true,
			FullTimestamp: true,
		})
		logrus.SetOutput(mw)
		logLevel := utils.GetEnvOrDefault("LOG_LEVEL", "info")

		// Parse log level string to logrus constant
		level, err := logrus.ParseLevel(logLevel)
		if err != nil {
			logrus.SetLevel(level)
		}
	})
}

func CloseLogFile() {
	if logFile != nil {
		logFile.Close()
	}
}

func ensureLogsDir(logsDir string) {
	_, err := os.Stat(logsDir)
	if err != nil {
		logrus.Fatal(err)
	}
	if os.IsNotExist(err) {
		err := os.MkdirAll(logsDir, 0755)
		if err != nil {
			logrus.Fatal(err)
		}
	}
}
