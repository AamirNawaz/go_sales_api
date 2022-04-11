package controllers

import (
	db "go_sales_api/config"
	"go_sales_api/middleware"
	"go_sales_api/models"
	"github.com/gofiber/fiber/v2"
	"log"
	"strconv"
	"strings"
)

func GetRevenues(c *fiber.Ctx) error {

	//Token authenticate
	headerToken := c.Get("Authorization")
	if headerToken == "" {
		return c.Status(404).JSON(fiber.Map{
			"status":  "error",
			"message": "Order does not exist",
		})
	}
	if err := middleware.AuthenticateToken(middleware.SplitToken(headerToken)); err != nil {
		return c.Status(404).JSON(fiber.Map{
			"status":  "error",
			"message": "Token expired or invalid",
		})
	}
	//Token authenticate

	order := []models.Order{}

	db.DB.Find(&order)
	TotalRevenues := make([]*models.RevenueResponse, 0)

	Resp1 := models.RevenueResponse{}
	Resp2 := models.RevenueResponse{}
	Resp3 := models.RevenueResponse{}

	sum1 := 0
	sum2 := 0
	sum3 := 0
	for _, v := range order {
		if v.PaymentTypesId == 1 {
			payment := models.Payment{}
			paymentTypes := models.PaymentType{}

			db.DB.Where("id=?", 1).First(&paymentTypes)
			db.DB.Where("payment_type_id=?", 1).First(&payment)

			sum1 += v.TotalPaid
			Resp1.Name = paymentTypes.Name
			Resp1.Logo = payment.Logo
			Resp1.TotalAmount = sum1
			Resp1.PaymentTypeId = v.PaymentTypesId
		}

		if v.PaymentTypesId == 2 {

			payment := models.Payment{}
			paymentTypes := models.PaymentType{}

			db.DB.Where("id=?", 2).First(&paymentTypes)
			db.DB.Where("payment_type_id=?", 2).First(&payment)

			sum2 += v.TotalPaid
			Resp2.Name = paymentTypes.Name
			Resp2.Logo = payment.Logo
			Resp2.TotalAmount = sum2
			Resp2.PaymentTypeId = v.PaymentTypesId
		}
		if v.PaymentTypesId == 3 {

			payment := models.Payment{}
			paymentTypes := models.PaymentType{}

			db.DB.Where("id=?", 3).First(&paymentTypes)
			db.DB.Where("payment_type_id=?", 3).First(&payment)

			sum3 += v.TotalPaid
			Resp3.Name = paymentTypes.Name
			Resp3.Logo = payment.Logo
			Resp3.TotalAmount = sum2
			Resp3.PaymentTypeId = v.PaymentTypesId
		}
	}
	TotalRevenues = append(TotalRevenues, &Resp1)
	TotalRevenues = append(TotalRevenues, &Resp2)
	TotalRevenues = append(TotalRevenues, &Resp3)

	return c.Status(200).JSON(fiber.Map{
		"success": true,
		"message": "success",
		"data": map[string]interface{}{
			"totalRevenue": sum1 + sum2 + sum3,
			"paymentTypes": TotalRevenues,
		},
	})
}

type Sold struct {
	ProductId   string `json:"productId"`
	Quantities  string `json:"quantities"`
	TotalAmount int    `json:"totalAmount"`
}

func GetSolds(c *fiber.Ctx) error {
	orders := []models.Order{}
	db.DB.Find(&orders)

	TotalSolds := make([]*models.SoldResponse, 0)
	TotalSoldsFinal := make([]*models.SoldResponse, 0)

	for _, v := range orders {
		quantities := strings.Split(v.Quantities, ",")
		quantities = quantities[1:]

		products := strings.Split(v.ProductId, ",")
		products = products[1:]

		for i := 0; i < len(products); i++ {
			prods := models.Product{}
			p, err := strconv.Atoi(products[i])
			q, errq := strconv.Atoi(quantities[i])

			if err != nil {
				log.Fatalf("->", err)
			}
			if errq != nil {
				log.Fatalf("->", errq)
			}

			db.DB.Where("id", p).Find(&prods)
			TotalSolds = append(TotalSolds, &models.SoldResponse{
				Name:        prods.Name,
				ProductId:   p,
				TotalQty:    q,
				TotalAmount: q * prods.Price,
			})
		}

	}
	duplicates := []int{}
	for _, v := range TotalSolds {

		if contains(duplicates, v.ProductId) == false {
			duplicates = append(duplicates, v.ProductId)
		}
	}
	quantityArray := []int{}
	for _, v := range duplicates {
		qty := 0
		for _, x := range TotalSolds {
			if v == x.ProductId {
				qty = qty + x.TotalQty
			}
		}
		quantityArray = append(quantityArray, qty)

	}

	for i := 0; i < len(duplicates); i++ {

		prods := models.Product{}
		db.DB.Where("id", duplicates[i]).Find(&prods)
		TotalSoldsFinal = append(TotalSoldsFinal, &models.SoldResponse{
			Name:        prods.Name,
			TotalQty:    quantityArray[i],
			TotalAmount: quantityArray[i] * prods.Price,
			ProductId:   duplicates[i],
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"success": true,
		"message": "success",
		"data": map[string]interface{}{
			"orderProducts": TotalSoldsFinal,
		},
	})
}

func contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
