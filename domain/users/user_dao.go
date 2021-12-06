package users

import (
	"log"
	"strings"

	"github.com/ezealcor/PointApp_users-api/utils/errors"
	"github.com/ezealcor/PointApp_users-api/utils/mysql_helpers"
)

const (
	QUERY_INSERT_USER     = "INSERT INTO users(first_name, last_name, email, created_date, user_type) VALUES (?, ?, ?, ?, ?)"
	QUERY_GET_USER        = "SELECT id, first_name, last_name, email, created_date, user_type FROM users WHERE id = ?"
	QUERY_UPDATE_USER     = "UPDATE users SET first_name = ?, last_name = ?, email = ?, user_type = ? WHERE id = ?"
	QUERY_EMAIL_EXISTS    = "SELECT email FROM users WHERE email = ?"
	EMAIL_UNIQUE          = "email_UNIQUE"
	NO_ROWS_IN_RESULT_SET = " no rows in result set"
)

func (user *User) Get() *errors.ResponseError {
	stmt, stmtError := mysql_helpers.PrepareStatement(QUERY_UPDATE_USER)
	if stmtError != nil {
		return stmtError
	}
	defer stmt.Close()

	getResult := stmt.QueryRow(user.Id)
	err := getResult.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.UserType, &user.CreatedDate)
	if err != nil {
		log.Println(err.Error())
		if strings.Contains(err.Error(), NO_ROWS_IN_RESULT_SET) {
			return errors.NotFoundError("user not found")
		}
		errors.InternalServerError("database error")
	}

	return nil
}

func (user *User) Save() *errors.ResponseError {

	stmt, stmtError := mysql_helpers.PrepareStatement(QUERY_INSERT_USER)
	if stmtError != nil {
		return stmtError
	}

	defer stmt.Close()

	insertResult, insertError := stmt.Exec(user.FirstName, user.LastName, user.Email, user.CreatedDate, user.UserType)

	if insertError != nil {
		log.Println(insertError.Error())
		if strings.Contains(insertError.Error(), EMAIL_UNIQUE) {
			return errors.BadRequestError("email already exists")
		}
		return errors.InternalServerError("database error")
	}

	userId, getIdError := insertResult.LastInsertId()
	if getIdError != nil {
		log.Println(getIdError.Error())
		return errors.InternalServerError("database error")
	}
	user.Id = userId

	return nil
}

func (user *User) Update() *errors.ResponseError {

	stmt, stmtError := mysql_helpers.PrepareStatement(QUERY_INSERT_USER)
	if stmtError != nil {
		return stmtError
	}

	defer stmt.Close()

	_, updateError := stmt.Exec(user.FirstName, user.LastName, user.Email, user.UserType, user.Id)

	if updateError != nil {
		log.Println(updateError.Error())
		if strings.Contains(updateError.Error(), EMAIL_UNIQUE) {
			return errors.BadRequestError("email already exists")
		}
		return errors.InternalServerError("database error")
	}

	return nil
}
