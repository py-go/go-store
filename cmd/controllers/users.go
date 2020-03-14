package controllers

import (
	"errors"
	"go-store/cmd/dtos"
	svc "go-store/cmd/services"

	"github.com/gin-gonic/gin"

	"go-store/cmd/models"

	"net/http"
)

func Signup(c *gin.Context) {

	var json dtos.Auth
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, dtos.ValidationErrorResponse(err))
		return
	}

	user := models.User{Username: json.Username}
	user.SetPassword(string(json.Password))
	if err := svc.CreateOne(&user); err != nil {
		c.JSON(http.StatusUnprocessableEntity, dtos.ErrorResponse("message", err))
		return
	}
	c.JSON(http.StatusCreated, dtos.UserDetails(&user))

}

func Login(c *gin.Context) {

	var json dtos.Auth
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, dtos.ValidationErrorResponse(err))
		return
	}

	user, err := svc.FindOneUser(&models.User{Username: json.Username})
	if err != nil {
		c.JSON(http.StatusBadRequest, dtos.ErrorResponse("username", err))
		return
	}

	if user.IsValidPassword(json.Password) != nil {
		c.JSON(http.StatusBadRequest, dtos.ErrorResponse("password", errors.New("invalid credentials")))
		return
	}
	c.JSON(http.StatusOK, dtos.UserDetails(&user))

}
