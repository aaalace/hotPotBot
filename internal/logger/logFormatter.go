package logger

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"time"
)

type LogFormatter struct {
	Location *time.Location
}

func NewLogFormatter(loc *time.Location) *LogFormatter {
	return &LogFormatter{Location: loc}
}

func (f *LogFormatter) Format(entry *logrus.Entry) ([]byte, error) {

	username, okUsername := entry.Data["username"]
	tgId, okTgId := entry.Data["tgId"]
	if okUsername && okTgId {
		return []byte(
			fmt.Sprintf("%s [%s] %s --- {%s, %v}\n",
				entry.Time.In(f.Location).Format("2006-01-02 15:04:05"),
				entry.Level,
				entry.Message,
				username,
				tgId,
			),
		), nil
	}
	return []byte(
		fmt.Sprintf("%s [%s] %s\n",
			entry.Time.In(f.Location).Format("2006-01-02 15:04:05"),
			entry.Level,
			entry.Message,
		),
	), nil
}
