package db

import (
	"errors"

	"log"
	dbConf "pos/internal/config/db"
	"pos/pkg/shared/enum"
	"sync"

	_ "github.com/lib/pq"
)

var (
	errorInvalidDbInstance = errors.New("Invalid db Instance")
)

var atomicinz uint64
var instanceDb map[string]DbDriver = make(map[string]DbDriver)

var once sync.Once

// DbDriver is object DB
type DbDriver interface {
	Db() interface{}
}

func NewInstanceDb(config dbConf.Database) (DbDriver, error) {
	var err error
	var dbName = config.Dbname

	once.Do(func() {
		switch config.Adapter {
		case enum.Postgres:
			dbConn, sqlErr := NewPostgresDriver(config)
			if sqlErr != nil {
				err = sqlErr
				log.Fatalf("Disconnect Database %v", err)
			}
			instanceDb[dbName] = dbConn
		default:
			err = errorInvalidDbInstance
		}
	})
	return instanceDb[dbName], err
}
