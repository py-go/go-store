package dtos

import (
	"go-store/cmd/models"
)

type SaleOrder struct {
	FirstName      string `form:"first_name" json:"first_name" xml:"first_name"`
	LastName       string `form:"last_name" json:"last_name" xml:"last_name"`
	Country        string `form:"country" json:"country" xml:"country"`
	City           string `form:"city" json:"city" xml:"city"`
	Street         string `form:"street" json:"street" xml:"street" `
	ZipCode        string `form:"zip_code" json:"zip_code" xml:"zip_code" `
	AddressId      uint   `form:_id" json:"address_id" xml:"address_id" `
	TrackingNumber string `form:"tracking_number" json:"tracking_number" xml:"tracking_number" `
	// AmountTotal float64 `form:"amount_total" json:"amount_total" xml:"amount_total" `
	// Discount    float64 `form:"discount" json:"discount" xml:"discount" `
	// Amount      float64 `form:"amount" json:"amount" xml:"amount" `

	OrderItems []struct {
		Id       uint `form:"id" json:"id" binding:"required"`
		Quantity int  `form:"quantity" json:"quantity" binding:"required"`
	} `form:"items" json:"items" xml:"items"  binding:"required"`
}

func CreateSaleOrder(order *models.Order, includes ...bool) map[string]interface{} {

	response := map[string]interface{}{
		"id":              order.ID,
		"tracking_number": order.TrackingNumber(),
		"state":           order.Status(),
		"amount_total":    order.AmountTotal,
		"amount":          order.Amount,
	}
	orderItems := make([]map[string]interface{}, len(order.OrderItems))
	for i := 0; i < len(order.OrderItems); i++ {
		item := order.OrderItems[i]
		orderItems[i] = map[string]interface{}{
			"name":  item.ProductName,
			"slug":  item.Slug,
			"price": item.Price,
			"qty":   item.Quantity,
		}
	}
	response["items"] = orderItems
	return response
}

func OrderData(order *models.Order) map[string]interface{} {
	return CreateSaleOrder(order, true, true, false)
}
