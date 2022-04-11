package models

import "time"

type Product struct {
	Id               int       `json:"id" gorm:"type:INT(10) UNSIGNED NOT NULL AUTO_INCREMENT;primaryKey"`
	Sku              string    `json:"sku"`
	Name             string    `json:"name"`
	Stock            int       `json:"stock"`
	Price            int       `json:"price"`
	Image            string    `json:"image"`
	TotalFinalPrice  int       `json:"total_final_price"`
	TotalNormalPrice int       `json:"total_normal_price"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
	CategoryId       int       `json:"categoryId"`
	DiscountId       int       `json:"discountId"`
}

type ProductResult struct {
	Id       int      `json:"productId" gorm:"type:INT(10) UNSIGNED NOT NULL AUTO_INCREMENT;primaryKey"`
	Sku      string   `json:"sku"`
	Name     string   `json:"name"`
	Stock    int      `json:"stock"`
	Price    int      `json:"price"`
	Image    string   `json:"image"`
	Category Category `json:"category"`
	Discount Discount `json:"discount"`
}
