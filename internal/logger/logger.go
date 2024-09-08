package logger

import (
	"github.com/sirupsen/logrus"
	"os"
	"time"
)

var Log *logrus.Logger

func init() {
	Log = logrus.New()

	// file work
	date := time.Now().Format("01-02-2006")
	file, err := os.OpenFile("logs/"+date, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		Log.Fatalf("Failed to open log file | %v", err.Error())
	}
	Log.SetOutput(file)

	Log.SetFormatter(&logrus.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: "01-02-2006 15:04:05",
	})
}
