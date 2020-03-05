package product

import (
	"errors"
	"go-store/cmd/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func AddProduct(router *gin.Engine) {
	router.GET("/", getProducts)
	router.GET("/products", getProducts)
	router.GET("/product/:product_id", getProduct)

}

func getProducts(c *gin.Context) {
	products := getAllProducts()
	utils.Render(c, "products.html", gin.H{
		"title":   "Products",
		"payload": products,
	})

}

func getProduct(c *gin.Context) {
	if productID, err := strconv.Atoi(c.Param("product_id")); err == nil {
		if product, err := getProductByID(productID); err == nil {
			utils.Render(c, "product.html", gin.H{
				"title":   product.Title,
				"payload": product,
			})

		} else {
			c.AbortWithError(http.StatusNotFound, err)
			utils.RenderError(c, http.StatusNotFound, "error.html", gin.H{
				"title":   "Error",
				"message": "Product not found",
			})
		}

	} else {
		c.AbortWithStatus(http.StatusNotFound)
		utils.RenderError(c, http.StatusNotFound, "error.html", gin.H{
			"message": "Product not found",
			"title":   "Error",
		})
	}
}

func getProductByID(id int) (*product, error) {
	for _, a := range productList {
		if a.ID == id {
			return &a, nil
		}
	}
	return nil, errors.New("product not found")
}
