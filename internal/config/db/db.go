package db

// DatabaseList is struct for list of database
type DatabaseList struct {
	HC struct {
		Postgres Database
	}
}

// Database is struct for database conf
type Database struct {
	Host     string
	Port     string
	Username string
	Password string
	Dbname   string
	Adapter  string
}
