package services

import (
	"github.com/ezealcor/PointApp_users-api/domain/users"
	"github.com/ezealcor/PointApp_users-api/utils/dates"
	"github.com/ezealcor/PointApp_users-api/utils/errors"
)

func CreateUser(user *users.User) (*users.User, *errors.ResponseError) {
	err := user.Validate()
	if err != nil {
		return nil, err
	}
	user.CreatedDate = dates.GetUTCString()
	errOnSave := user.Save()

	if errOnSave != nil {
		return nil, errOnSave
	}
	return user, nil
}

func GetUser(user *users.User) (*users.User, *errors.ResponseError) {
	if user.Id <= 0 {
		return nil, errors.BadRequestError("invalid user id")
	}
	err := user.Get()
	if err != nil {
		return nil, err
	}
	return user, nil

}

func UpdateUser(user *users.User) (*users.User, *errors.ResponseError) {
	err := user.Validate()
	if err != nil {
		return nil, err
	}

	errOnSave := user.Update()

	if errOnSave != nil {
		return nil, errOnSave
	}
	return user, nil
}
