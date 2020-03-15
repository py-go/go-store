package controllers

import (
	"errors"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"

	"go-store/cmd/dtos"
	"go-store/cmd/models"
	svc "go-store/cmd/services"
	"net/http"
)

func ListOrders(c *gin.Context) {
	user := c.MustGet("storeUser").(models.User)
	obj, err := svc.FetchOrdersByUser(user)

	if err != nil {
		dtos.RenderOrmError(c, err)
		return
	}
	c.JSON(http.StatusOK, obj)
}

func GetOrder(c *gin.Context) {
	objID, err := strconv.Atoi(c.Param("id"))
	user := c.MustGet("storeUser").(models.User)
	obj, err := svc.FetchOrderDetails(uint(objID))

	if err != nil {
		dtos.RenderOrmError(c, err)
		return
	}
	if obj.UserID != user.ID {
		c.AbortWithStatusJSON(http.StatusForbidden,
			dtos.ErrorResponse("message", errors.New("Insufficient permission.")))
		return
	}
	c.JSON(http.StatusOK, dtos.OrderData(&obj))
}

func CreateOrder(c *gin.Context) {
	var json dtos.SaleOrder

	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, dtos.ValidationErrorResponse(err))
		return
	}

	order := models.Order{
		OrderStatus: 0,
	}
	var user models.User
	userObj, userLoggedIn := c.Get("storeUser")
	if userLoggedIn {
		user = (userObj).(models.User)
		order.User = user
	}

	err, status := AddToCart(&order, json, true)
	if err != nil {
		c.AbortWithStatusJSON(status, dtos.ErrorResponse("message", err))
		return
	}

	err = svc.CreateOne(&order)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, dtos.ErrorResponse("message", err))
		return
	}
	c.JSON(http.StatusOK, dtos.OrderData(&order))

}

func AddToCart(order *models.Order, json dtos.SaleOrder, newCart bool) (err error, status int) {

	cartItemCount := len(json.OrderItems)

	var productIds = make([]uint, cartItemCount)
	var itemQty = make(map[uint]int)
	for index, item := range json.OrderItems {
		productIds[index] = item.Id
		if _, ok := itemQty[item.Id]; ok {
			itemQty[item.Id] += item.Quantity
		} else {
			itemQty[item.Id] = item.Quantity
		}

	}

	products, err := svc.FechProductByIDs(productIds, []string{"id", "name", "slug", "price"})
	if err != nil {
		status = http.StatusNotFound
		return
	}
	productCount := len(products)
	if newCart {
		order.OrderItems = make([]models.OrderItem, productCount)
	}

	AmountTotal := 0.0
	var itemPrice = make(map[uint]float64)
	for index, product := range products {
		qty := itemQty[product.ID]
		price := product.Price * float64(qty)
		itemPrice[product.ID] = product.Price
		order.OrderItems[index] = models.OrderItem{
			ProductID:   product.ID,
			ProductName: product.Name,
			Slug:        product.Slug,
			Price:       product.Price,
			Quantity:    qty,
		}
		AmountTotal += price
		log.Println(price, AmountTotal, "SSSSSSSSSSSs")
	}

	order.AmountTotal = (AmountTotal)
	cartDiscount, err := computeCartRuleDiscount(order, itemQty, itemPrice)
	if err != nil {
		status = http.StatusNotFound
		return
	}
	order.Discount = cartDiscount
	// ADD Coupon Discount logic

	return
}

func computeCartRuleDiscount(order *models.Order, itemQty map[uint]int, itemPrice map[uint]float64) (result float64, err error) {
	discounts, err := svc.FetchProductCartRule(order.User)
	if err != nil {
		log.Println("ERROR:orders.computeDiscount:", err)
	}
	for _, discount := range discounts {
		percerange := discount.Percerange
		combination := discount.Type
		if combination != len(discount.Rules) {
			continue
		}
		productsFactor := make(map[uint]int)
		minFactor := 0
		for index, rule := range discount.Rules {
			dQty := itemQty[rule.ProductID] / rule.Quantity
			if dQty == 0 {
				productsFactor = make(map[uint]int)
				break
			}
			if index == 0 {
				minFactor = dQty
			}
			if minFactor > dQty {
				minFactor = dQty
			}
			productsFactor[rule.ProductID] = dQty
		}
		for productID, factor := range productsFactor {
			price := itemPrice[productID]
			discount := (price * percerange / float64(100))
			result += discount
			log.Println(percerange, productID, price, factor, minFactor, discount, "============2222222222222222")
		}

	}
	return
}
