package context

import (
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/jmoiron/sqlx"
	"sync"
)

type AppContext struct {
	DB           *sqlx.DB
	S3Client     *s3.S3
	UserRequests map[int64]string // временное решение
	Mu           sync.Mutex
}
