package controllers

import (
	db "go_sales_api/config"
	"go_sales_api/models"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"os"
	"strconv"
	"time"
)

func Login(c *fiber.Ctx) error {
	cashierId := c.Params("cashierId")
	var data map[string]string
	if err := c.BodyParser(&data); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"success": false,
			"Message": "Invalid post request",
		})
	}

	//check if passcode is empty
	if data["passcode"] == "" {
		return c.Status(400).JSON(fiber.Map{
			"success": false,
			"message": "Passcode is required",
			"error":   map[string]interface{}{},
		})
	}
	var cashier models.Cashier
	db.DB.Where("id = ?", cashierId).First(&cashier)

	//check if cashier exist
	if cashier.Id == 0 {
		return c.Status(404).JSON(fiber.Map{
			"success": false,
			"message": "Cashier Not found",
			"error":   map[string]interface{}{},
		})
	}
	//check if passcode match
	//fmt.Println("--------------------------------")
	//fmt.Println("--------------DB Passcode------------------", cashier.Passcode)
	//fmt.Println("--------------DB Passcode typeOf------------------", reflect.TypeOf(cashier.Passcode))
	//fmt.Println("--------------------------------")
	//
	//fmt.Println("--------------------------------")
	//fmt.Println("--------------body passcode------------------", data["passcode"])
	//fmt.Println("--------------DB Passcode typeOf------------------", reflect.TypeOf(data["passcode"]))
	//fmt.Println("--------------------------------")

	if cashier.Passcode != data["passcode"] {
		return c.Status(401).JSON(fiber.Map{
			"success": false,
			"message": "Passcode Not Match",
			"error":   map[string]interface{}{},
		})
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"Issuer":    strconv.Itoa(int(cashier.Id)),
		"ExpiresAt": time.Now().Add(time.Hour * 24).Unix(), //1 day
	})
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return c.Status(401).JSON(fiber.Map{
			"success": false,
			"message": "Token Expired or invalid",
		})
	}

	cashierData := make(map[string]interface{})
	cashierData["token"] = tokenString

	return c.Status(200).JSON(fiber.Map{
		"success": true,
		"message": "Success",
		"data":    cashierData,
	})

}
func Logout(c *fiber.Ctx) error {
	cashierId := c.Params("cashierId")
	var data map[string]string
	if err := c.BodyParser(&data); err != nil {
		return err
	}
	//check if passcode is empty
	if data["passcode"] == "" {
		return c.Status(400).JSON(fiber.Map{
			"success": false,
			"Message": "Passcode is required",
		})
	}

	var cashier models.Cashier
	db.DB.Where("Id = ?", cashierId).First(&cashier)

	//check if cashier exist
	if cashier.Id == 0 {
		return c.Status(404).JSON(fiber.Map{
			"success": false,
			"Message": "Cashier Not found",
		})
	}
	//check if passcode match
	if cashier.Passcode != data["passcode"] {
		return c.Status(401).JSON(fiber.Map{
			"success": false,
			"Message": "Passcode Not Match",
		})
	}
	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HTTPOnly: true,
	}

	c.Cookie(&cookie)
	return c.Status(200).JSON(fiber.Map{
		"success": true,
		"Message": "success",
	})
}
func Passcode(c *fiber.Ctx) error {
	cashierId := c.Params("cashierId")
	var cashier models.Cashier
	db.DB.Select("id,name,passcode").Where("id=?", cashierId).First(&cashier)

	if cashier.Name == "" || cashier.Id == 0 {
		return c.Status(404).JSON(fiber.Map{
			"success": false,
			"message": "Cashier Not Found",
			"error":   map[string]interface{}{},
		})
	}

	cashierData := make(map[string]interface{})
	cashierData["passcode"] = cashier.Passcode

	return c.Status(200).JSON(fiber.Map{
		"success": true,
		"message": "Success",
		"data":    cashierData,
	})
}
