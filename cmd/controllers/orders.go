package controllers

import (
	"errors"
	"fmt"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"

	"go-store/cmd/dtos"
	"go-store/cmd/models"
	svc "go-store/cmd/services"
	"net/http"
)

//https://github.com/cooldryplace/cart/blob/master/cart.go
//
// type storage interface {
// 	AddProduct(ctx context.Context, cartID, productID int64, quantity uint32) error
// 	DeleteProduct(ctx context.Context, cartID, productID int64) error
// 	CartByID(ctx context.Context, id int64) (Cart, error)
// 	CreateCart(ctx context.Context, cart Cart) (Cart, error)
// 	DeleteCart(ctx context.Context, cartID int64) error
// 	DeleteLineItems(ctx context.Context, cartID int64) error
// }

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
	if obj.UserId != user.ID {
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
	log.Println(userObj, userLoggedIn)
	// os.Exit(0)
	if userLoggedIn {
		user = (userObj).(models.User)
		// order.UserId = user.ID
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

// func MergeCartItem(newItem *CartItem, order *models.Order) {
// 	exists := false
// 	for idx, item := range order.Items {
// 		if item.Name == newItem.Name {
// 			order.Items[idx].Qty += newItem.Qty // Increment qty
// 			exists = true
// 		}
// 	}
// 	if !exists {
// 		order.Items = append(order.Items, *newItem)
// 	}
// 	order.Total += newItem.Price
// }

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
	// for i := 0; i < cartItemCount; i++ {
	// 	productId := json.OrderItems[i].Id
	// 	productIds[i] = productId
	// 	if _, ok := itemQty[productId]; ok {
	// 		itemQty[productId] += json.OrderItems[i].Quantity
	// 	} else {
	// 		itemQty[productId] = json.OrderItems[i].Quantity
	// 	}

	// }

	products, err := svc.FechProductByIDs(productIds, []string{"id", "name", "slug", "price"})
	if err != nil {
		status = http.StatusNotFound
		return
	}
	productCount := len(products)
	log.Println(productIds, products, productCount, "SSSSSSSSS")
	log.Println(fmt.Sprintf("%#v", itemQty), "SSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSS")
	// os.Exit(0)

	// if productCount != len(itemQty) {
	// 	err = errors.New("products are missing")
	// 	status = http.StatusGone
	// 	return
	// }
	if newCart {
		order.OrderItems = make([]models.OrderItem, productCount)
	}
	AmountTotal := 0
	for index, product := range products {
		qty := itemQty[product.ID]
		price := product.Price * qty
		order.OrderItems[index] = models.OrderItem{
			ProductId:   product.ID,
			ProductName: product.Name,
			Slug:        product.Slug,
			Price:       product.Price,
			Quantity:    qty,
		}
		AmountTotal += price
	}
	// for productId, qty := range itemQty {

	// }
	//Discount := computeDiscount(Discount, Amount)
	//
	log.Println(fmt.Sprintf("%v=============", order.OrderItems))
	log.Println(len(order.OrderItems), "@@@@@@@@@@@@@@@@@")

	// os.Exit(0)
	return
}

// func UpdateCart(order *models.Order, json dtos.SaleOrder) (err error, status int) {

// 	cartItemCount := len(json.OrderItems)

// 	var productIds = make([]uint, cartItemCount)
// 	for i := 0; i < cartItemCount; i++ {
// 		productIds[i] = json.OrderItems[i].Id
// 	}

// 	products, err := svc.FechProductByIDs(productIds, []string{"id", "name", "slug", "price"})
// 	if err != nil {
// 		status = http.StatusNotFound
// 		return
// 	}

// 	productCount := len(products)
// 	orderItems := make([]models.OrderItem, productCount)

// 	for i := 0; i < productCount; i++ {
// 		orderItems[i] = models.OrderItem{
// 			ProductId:   products[i].ID,
// 			ProductName: products[i].Name,
// 			Slug:        products[i].Slug,
// 			Quantity:    json.OrderItems[i].Quantity,
// 			Price:       products[i].Price,
// 		}
// 	}
// 	order.OrderItems = orderItems
// 	return
// }
// func AddItem(newItem *CartItem) {
// 	exists := false
// 	for idx, item := range ShoppingCart.Items {
// 		if item.Name == newItem.Name {
// 			ShoppingCart.Items[idx].Qty += newItem.Qty // Increment qty
// 			exists = true
// 		}
// 	}
// 	if !exists {
// 		ShoppingCart.Items = append(ShoppingCart.Items, *newItem)
// 	}
// 	ShoppingCart.Total += newItem.Price
// }
