package db

import (
	"database/sql"
	"log"
	dbConf "pos/internal/config/db"
)

type DriverPostgres struct {
	config dbConf.Database
	db     *sql.DB
}

// NewPostgressDriver new object SQL Driver
func NewPostgresDriver(config dbConf.Database) (DbDriver, error) {
	dbConn, err := connect(config)
	if err != nil {
		return nil, err
	}

	return &DriverPostgres{
		config: config,
		db:     dbConn,
	}, nil
}

func connect(config dbConf.Database) (*sql.DB, error) {
	user := config.Username
	password := config.Password
	// host := config.Host
	dbname := config.Dbname

	dbConn, err := sql.Open("postgres", "user="+user+" password="+password+" dbname="+dbname+" sslmode=disable")
	if err != nil {
		log.Fatalln(err.Error())
		return nil, err
	}
	return dbConn, err
}

// Db get db instance of sql
func (m *DriverPostgres) Db() interface{} {
	return m.db
}
