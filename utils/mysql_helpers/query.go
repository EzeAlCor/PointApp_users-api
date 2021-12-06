package mysql_helpers

import (
	"database/sql"
	"log"

	"github.com/ezealcor/PointApp_users-api/data_sources/mysql/users_db"
	"github.com/ezealcor/PointApp_users-api/utils/errors"
)

func PrepareStatement(query string) (*sql.Stmt, *errors.ResponseError) {
	stmt, stmtError := users_db.Client.Prepare(query)
	if stmtError != nil {
		log.Println("\033[32m", stmtError.Error())
		return nil, errors.InternalServerError("invalid sql statement")
	}
	return stmt, nil
}
