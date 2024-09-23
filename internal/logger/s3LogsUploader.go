package logger

import (
	"hotPotBot/internal/consts"
	"os"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

type S3LogsUploader struct {
	bucket        string
	keyPrefix     string
	uploader      *s3manager.Uploader
	location      *time.Location
	localLogsPath string
}

func NewS3LogsUploader(loc *time.Location, localLogsPath string) *S3LogsUploader {
	sess := session.Must(session.NewSession(&aws.Config{
		Region: aws.String("eu-central-1"),
		Credentials: credentials.NewStaticCredentials(
			os.Getenv("S3_ACCESS_KEY"),
			os.Getenv("S3_SECRET_KEY"),
			""),
		Endpoint: aws.String(os.Getenv("S3_PATH")),
	}))

	return &S3LogsUploader{
		bucket:        os.Getenv("S3_BUCKET"),
		keyPrefix:     "logs/",
		uploader:      s3manager.NewUploader(sess),
		location:      loc,
		localLogsPath: localLogsPath,
	}
}

func (u *S3LogsUploader) clearLocalLogs() error {
	file, err := os.OpenFile(u.localLogsPath, os.O_WRONLY|os.O_TRUNC, 0666)
	if err != nil {
		return err
	}
	defer file.Close()

	err = file.Truncate(0)
	if err != nil {
		return err
	}

	return nil
}

func (u *S3LogsUploader) flush() error {
	file, err := os.Open(u.localLogsPath)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = u.uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(u.bucket),
		Key:    aws.String(u.keyPrefix + time.Now().In(u.location).Format("2006-01-02") + ".log"),
		Body:   file,
	})
	if err != nil {
		return err
	}

	return nil
}

func (u *S3LogsUploader) Listen() {
	ticker := time.NewTicker(consts.LogsFlushInterval)
	Log.Info("TICKER GOROUTINE | Ticker started")
	for {
		select {
		case <-ticker.C:
			err := u.flush()
			if err != nil {
				Log.Errorf("TICKER GOROUTINE | Error in flushing logs to S3 | %v", err.Error())
			} else {
				err = u.clearLocalLogs()
				if err != nil {
					Log.Errorf("TICKER GOROUTINE | Error clearing local logs | %v", err.Error())
				} else {
					Log.Info("TICKER GOROUTINE | Previous logs correctly sent to S3")
				}
			}
		}
	}
}
