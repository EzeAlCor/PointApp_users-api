package users

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/ezealcor/PointApp_users-api/domain/users"
	"github.com/ezealcor/PointApp_users-api/services"
	"github.com/ezealcor/PointApp_users-api/utils/errors"
	"github.com/gin-gonic/gin"
)

func GetUser(c *gin.Context) {
	userId, parseErr := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if parseErr != nil {
		err := errors.BadRequestError("invalid user id")
		c.JSON(err.Status, err)
		return
	}
	user := users.User{Id: userId}
	result, err := services.GetUser(&user)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, result)
}

func CreateUser(c *gin.Context) {
	var user users.User
	err := c.ShouldBindJSON(&user)

	if err != nil {
		resError := errors.BadRequestError("invalid body request")
		c.JSON(resError.Status, resError)
		return
	}

	result, serviceError := services.CreateUser(&user)

	if serviceError != nil {
		c.JSON(serviceError.Status, serviceError)
		return
	}

	c.JSON(http.StatusCreated, result)

	fmt.Println("user", user)

}

func SearchtUser(c *gin.Context) {

}
