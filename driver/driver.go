package driver

import (
	"books-list/utils"
	"database/sql"
	"github.com/lib/pq"
	"os"
)

var db *sql.DB
func ConnectDB() *sql.DB {
	pgUrl, err := pq.ParseURL(os.Getenv("ELEPHANTSQL_URL"))
	utils.LogFatal(err)

	db, err = sql.Open("postgres", pgUrl)
	utils.LogFatal(err)

	err = db.Ping()
	utils.LogFatal(err)

	return db
}
