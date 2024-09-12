package s3

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"hotPotBot/internal/config"
	"hotPotBot/internal/logger"
	"io"
	"os"
)

func ConnectS3Client(cfg *config.Config) *s3.S3 {
	logger.Log.Info("Initializing S3")

	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String("eu-central-1"),
		Credentials: credentials.NewStaticCredentials(cfg.S3AccessKey, cfg.S3SecretKey, ""),
		Endpoint:    aws.String(cfg.S3Path),
	})
	if err != nil {
		logger.Log.Errorf("Failed to setup S3 | %v", err.Error())
		return nil
	}

	return s3.New(sess)
}

func DownloadImageFromS3(s3Client *s3.S3, key string) (io.ReadCloser, error) {
	input := &s3.GetObjectInput{
		Bucket: aws.String(os.Getenv("S3_BUCKET")),
		Key:    aws.String(key),
	}

	result, err := s3Client.GetObject(input)
	if err != nil {
		return nil, err
	}

	return result.Body, nil
}
