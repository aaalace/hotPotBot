package logger

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"hotPotBot/consts"
	"os"
	"time"
)

var Log *logrus.Logger

type CustomFormatter struct {
	Location *time.Location
}

func (f *CustomFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	localTime := entry.Time.In(f.Location)

	return []byte(
		fmt.Sprintf("%s [%s] %s\n",
			localTime.Format("2006-01-02 15:04:05"),
			entry.Level,
			entry.Message),
	), nil
}

func init() {
	loc, err := time.LoadLocation(consts.AppLocation)
	if err != nil {
		panic(err)
	}

	Log = logrus.New()

	Log.SetFormatter(&CustomFormatter{Location: loc})
	Log.SetOutput(os.Stdout)
}
