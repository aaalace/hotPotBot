package context

import (
	"sync"

	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/jmoiron/sqlx"
)

type AppContext struct {
	DB           *sqlx.DB
	S3Client     *s3.S3
	UserRequests map[int64]string // временное решение
	Mu           sync.Mutex
}
