package logger

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"hotPotBot/internal/consts"
	"os"
	"time"
)

var Log *logrus.Logger

func init() {
	if err := godotenv.Load(); err != nil {
		panic("Error loading env variables")
	}

	Log = NewLogger()
	Log.Info("Logger initialized")
}

func NewLogger() *logrus.Logger {
	logger := logrus.New()

	loc, err := time.LoadLocation(consts.AppLocation)
	if err != nil {
		panic(fmt.Sprintf("Can not load location in logger | %v", err))
	}
	logger.SetFormatter(NewLogFormatter(loc))

	logger.SetOutput(os.Stdout)
	if mode := os.Getenv("MODE"); mode == "prod" {
		localLogsPath := "daily.log"
		file, err := os.OpenFile(localLogsPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			panic(fmt.Sprintf("Can not open local logs path | %v", err))
		}
		logger.SetOutput(file)

		s3LogsUploader := NewS3LogsUploader(loc, localLogsPath)
		go s3LogsUploader.Listen()
	}

	return logger
}
