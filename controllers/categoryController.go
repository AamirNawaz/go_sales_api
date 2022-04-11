package controllers

import (
	db "go_sales_api/config"
	"go_sales_api/middleware"
	"go_sales_api/models"
	"github.com/gofiber/fiber/v2"
	"log"
	"strconv"
)

// Category struct with two values
type Category struct {
	Id   uint   `json:"categoryId"`
	Name string `json:"name"`
}

func CreateCategory(c *fiber.Ctx) error {
	var data map[string]string
	err := c.BodyParser(&data)
	if err != nil {
		log.Fatalf("category error in post request %v", err)
	}
	if data["name"] == "" {
		return c.Status(400).JSON(fiber.Map{
			"success": false,
			"message": "Category name is required",
			"error":   map[string]interface{}{},
		})
	}
	category := models.Category{
		Name: data["name"],
	}
	db.DB.Create(&category)
	//result:=db.DB.Create(&category)

	//if result.RowsAffected == 0 {
	//	return c.Status(404).JSON(fiber.Map{
	//		"success": false,
	//		"Message": "category insertion failed",
	//	})
	//}

	return c.Status(200).JSON(fiber.Map{
		"success": true,
		"Message": "Success",
		"data":    category,
	})
}

func GetCategoryDetails(c *fiber.Ctx) error {
	categoryId := c.Params("categoryId")

	//Token authenticate
	headerToken := c.Get("Authorization")
	if headerToken == "" {
		return c.Status(401).JSON(fiber.Map{
			"success": false,
			"message": "Unauthorized",
			"error":   map[string]interface{}{},
		})
	}
	if err := middleware.AuthenticateToken(middleware.SplitToken(headerToken)); err != nil {
		return c.Status(401).JSON(fiber.Map{
			"success": false,
			"message": "Unauthorized",
			"error":   map[string]interface{}{},
		})
	}
	//Token authenticate

	var category models.Category
	db.DB.Select("id ,name").Where("id=?", categoryId).First(&category)
	categoryData := make(map[string]interface{})
	categoryData["categoryId"] = category.Id
	categoryData["name"] = category.Name

	if category.Name == "" {
		return c.Status(404).JSON(fiber.Map{
			"success": false,
			"Message": "No category found",
		})
	}

	return c.JSON(fiber.Map{
		"success": true,
		"Message": "Success",
		"data":    categoryData,
	})
}

func DeleteCategory(c *fiber.Ctx) error {

	categoryId := c.Params("categoryId")
	var category models.Category
	db.DB.Where("id=?", categoryId).First(&category)

	if category.Id == 0 {
		return c.Status(404).JSON(fiber.Map{
			"success": false,
			"message": "category not found",
			"error":   map[string]interface{}{},
		})
	}

	db.DB.Where("id = ?", categoryId).Delete(&category)

	return c.Status(200).JSON(fiber.Map{
		"success": true,
		"Message": "Success",
	})
}

func UpdateCategory(c *fiber.Ctx) error {
	categoryId := c.Params("categoryId")
	var category models.Category

	db.DB.Find(&category, "id = ?", categoryId)

	if category.Name == "" {
		return c.Status(404).JSON(fiber.Map{
			"success": false,
			"message": "Category not exist against this id",
		})
	}

	var updateCashierData models.Category
	c.BodyParser(&updateCashierData)
	if updateCashierData.Name == "" {
		return c.Status(400).JSON(fiber.Map{
			"success": false,
			"message": "Category name is required",
			"error":   map[string]interface{}{},
		})
	}
	category.Name = updateCashierData.Name
	db.DB.Save(&category)
	return c.Status(200).JSON(fiber.Map{
		"success": true,
		"Message": "Success",
		"data":    category,
	})

}

func CategoryList(c *fiber.Ctx) error {
	//Token authenticate
	headerToken := c.Get("Authorization")
	if headerToken == "" {
		return c.Status(401).JSON(fiber.Map{
			"success": false,
			"message": "Unauthorized",
			"error":   map[string]interface{}{},
		})
	}
	if err := middleware.AuthenticateToken(middleware.SplitToken(headerToken)); err != nil {
		return c.Status(401).JSON(fiber.Map{
			"success": false,
			"message": "Unauthorized",
			"error":   map[string]interface{}{},
		})
	}
	//Token authenticate

	limit, _ := strconv.Atoi(c.Query("limit"))
	skip, _ := strconv.Atoi(c.Query("skip"))
	var count int64
	var category []Category
	db.DB.Select("id ,name").Limit(limit).Offset(skip).Find(&category).Count(&count)
	metaMap := map[string]interface{}{
		"total": count,
		"limit": limit,
		"skip":  skip,
	}
	categoriesData := map[string]interface{}{
		"categories": category,
		"meta":       metaMap,
	}

	return c.JSON(fiber.Map{
		"success": true,
		"Message": "Success",
		"data":    categoriesData,
	})

}

//
//func CatL(c *fiber.Ctx) error {
//	auth := c.Get("Authorization")
//	if auth == "" {
//		return c.Status(401).JSON(fiber.Map{
//			"success": false,
//			"message": "Unauthorized",
//			"error":   map[string]interface{}{},
//		})
//	}
//
//	if err := util.AuthToken(util.SplitToken(auth)); err != nil {
//		return c.Status(401).JSON(fiber.Map{
//			"status":  "error",
//			"message": "Token expired or invalid",
//			"error":   map[string]interface{}{},
//		})
//	}
//
//	limit, _ := strconv.Atoi(c.Query("limit"))
//	skip, _ := strconv.Atoi(c.Query("skip"))
//	var countes int64
//
//	var catagory []model.Category
//	db.DB.Select([]string{"category_id ,name"}).Limit(limit).Offset(skip).Find(&catagory).Count(&countes)
//
//	return c.Status(404).JSON(fiber.Map{
//		"success": true,
//		"message": "Success",
//		"data":    catagory,
//		"meta": map[string]interface{}{
//			"total": countes,
//			"limit": limit,
//			"skip":  skip,
//		},
//	})
//
//}
