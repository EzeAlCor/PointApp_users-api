package users_db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var (
	Client *sql.DB
	err    error
	user   = os.Getenv("udbu")
	pass   = os.Getenv("udbp")
	host   = os.Getenv("udba")
	schema = os.Getenv("udbs")
)

func init() {
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s",
		user,
		pass,
		host,
		schema,
	)

	Client, err = sql.Open("mysql", dataSourceName)
	if err != nil {
		panic(err)
	}

	pringErr := Client.Ping()
	if pringErr != nil {
		panic(pringErr)
	}

	log.Println("database connection successfully configured")
}
