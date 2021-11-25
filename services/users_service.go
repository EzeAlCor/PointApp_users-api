package services

import (
	"github.com/ezealcor/PointApp_users-api/domain/users"
	"github.com/ezealcor/PointApp_users-api/utils/errors"
)

func CreateUser(user *users.User) (*users.User, *errors.ResponseError) {
	err := user.Validate()
	if err != nil {
		return nil, err
	}
	errOnSave := user.Save()

	if errOnSave != nil {
		return nil, err
	}
	return user, nil
}
