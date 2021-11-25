package users

import (
	"regexp"
	"strings"

	"github.com/ezealcor/PointApp_users-api/utils/errors"
)

type User struct {
	Id          int64  `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	CreatedDate string `json:"create_date"`
	Email       string `json:"email"`
	UserType    string `json:"user_type"`
}

func (user *User) Validate() *errors.ResponseError {
	formatedEmail := strings.TrimSpace(strings.ToLower(user.Email))
	emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	if emailRegex.MatchString(formatedEmail) {
		return nil
	}
	return errors.BadRequestError("invalid email")
}
