package users

import (
	"github.com/ezealcor/PointApp_users-api/utils/errors"
)

var userDB = make(map[int64]*User)

func (user *User) Get() *errors.ResponseError {
	result, ok := userDB[int64((user.Id))]

	if !ok {
		return errors.NotFoundError("user not found")
	}

	user.Id = result.Id
	user.FirstName = result.FirstName
	user.LastName = result.LastName
	user.Email = result.Email
	user.UserType = result.UserType
	user.CreatedDate = result.CreatedDate

	return nil
}

func (user *User) Save() *errors.ResponseError {
	currentUser := userDB[int64((user.Id))]

	if currentUser != nil {
		return errors.BadRequestError("user already exist")
	}

	userDB[int64((user.Id))] = user

	return nil
}
