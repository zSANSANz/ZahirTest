package models

import (
	"time"
)

type PaymentMethod struct {
	PaymentMethodId	    string         `gorm:"type:varchar(64);unique;primaryKey" json:"payment_method_id" form:"payment_method_id"`
	MethodName 			string         `gorm:"type:varchar(70);not null" json:"method_name" form:"method_name"`
	Code				string         `gorm:"type:varchar(10);not null" json:"code" form:"code"`
	CreatedDate   		time.Time      `json:"created_date" form:"created_date"`
}
