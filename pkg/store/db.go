package store

import (
	"database/sql"

	_ "github.com/lib/pq"
)

// DB is the database connection
var DB *sql.DB

// InitDB initializes the database connection
func InitDB(dataSourceName string) error {
	var err error
	DB, err = sql.Open("postgres", dataSourceName)
	if err != nil {
		return err
	}
	return DB.Ping()
}
