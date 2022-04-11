package models

import "time"

type Discount struct {
	Id              int       `json:"id" gorm:"type:INT(10) UNSIGNED NOT NULL AUTO_INCREMENT;primaryKey"`
	Qty             int       `json:"qty"`
	Type            string    `json:"type"`
	Result          int       `json:"result"`
	ExpiredAt       int       `json:"expiredAt"`
	ExpiredAtFormat string    `json:"expiredAtFormat"`
	StringFormat    string    `json:"stringFormat"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}
