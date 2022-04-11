package models

import "time"

type Cashier struct {
	Id        uint      `json:"id" gorm:"type:INT(10) UNSIGNED NOT NULL AUTO_INCREMENT;primaryKey"`
	Name      string    `json:"name"`
	Passcode  string    `json:"passcode"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
