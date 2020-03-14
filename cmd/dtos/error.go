package dtos

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"gopkg.in/go-playground/validator.v9"
)

type Error struct {
	Base
	Errors map[string]interface{} `json:"errors"`
}

// type HttpError struct {
// 	status      int
// 	description string
// }

func ErrorResponse(key string, err error) Error {
	res := Error{}
	res.Errors = map[string]interface{}{key: err.Error()}
	return res
}

func ValidationErrorResponse(err error) Error {
	errs := err.(validator.ValidationErrors)

	res := Error{}
	res.Errors = make(map[string]interface{})
	res.Meta = make([]string, len(errs))

	count := 0
	for _, v := range errs {
		tag := v.ActualTag()
		field := v.Field()
		if tag == "required" {
			var message = fmt.Sprintf("%v is required", field)
			res.Errors[field] = message
			res.Meta[count] = message
		} else {
			var message = fmt.Sprintf("%v has to be %v", field, tag)
			res.Errors[field] = message
			res.Meta = append(res.Meta, message)
		}
		count++
	}
	return res
}
func RenderOrmError(c *gin.Context, err error) {
	if gorm.IsRecordNotFoundError(err) == true {
		c.AbortWithStatusJSON(http.StatusNotFound, ErrorResponse("id", err))

	} else {
		c.AbortWithStatusJSON(http.StatusInternalServerError, ErrorResponse("message", err))

	}
}
