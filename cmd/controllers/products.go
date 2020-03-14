package controllers

import (
	"strconv"

	"github.com/gin-gonic/gin"

	"go-store/cmd/dtos"
	svc "go-store/cmd/services"
	"net/http"
)

func ListProduct(c *gin.Context) {
	obj, err := svc.FetchProducts()
	if err != nil {
		dtos.RenderOrmError(c, err)
		return
	}
	c.JSON(http.StatusOK, obj)
}

func GetProduct(c *gin.Context) {
	objID, err := strconv.Atoi(c.Param("id"))
	//c.MustGet("storeUser")
	obj, err := svc.FechProductByID(uint(objID))
	if err != nil {
		dtos.RenderOrmError(c, err)
		return
	}
	c.JSON(http.StatusOK, obj)
}
