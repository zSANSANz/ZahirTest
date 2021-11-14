package models

import (
	"time"
)

type OrderDetail struct {
	OrderDetailId      	string         `gorm:"type:varchar(64);primaryKey" json:"order_detail_id" form:"order_detail_id"`
	OrderId 			string         `gorm:"type:varchar(64);not null" json:"company_name" form:"company_name"`
	Order   			Order	   	   `json:"order" form:"order"`
	ProductId			string         `gorm:"type:varchar(64);unique;not null" json:"company_name" form:"company_name"`
	Product   			Product	   	   `json:"product" form:"product"`
	Qty					uint		   `json:"created_at" form:"created_at"`
	CreatedDate			time.Time      `json:"created_at" form:"created_at"`
}
