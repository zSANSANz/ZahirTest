package models

import (
	"time"
)

type Order struct {
	OrderId	        	string         `gorm:"type:varchar(64);unique;primaryKey" json:"order_id" form:"order_id"`
	CustomerId	        string         `gorm:"type:varchar(64);not_null" json:"customer_id" form:"customer_id"`
	Customer   			Customer	   `json:"customer" form:"customer"`
	OrderNumber			string         `gorm:"type:varchar(40);not null" json:"order_number" form:"order_number"`
	OrderDate  			time.Time      `gorm:"not null" json:"order_date" form:"order_date"`
	PaymentMethodId		string		   `gorm:"type:varchar(64);not_null" json:"payment_method_id" form:"payment_method_id"`
	PaymentMethod   	PaymentMethod  `json:"payment_method" form:"payment_method"`
}
