package models

import (
	"time"
)

type Customer struct {
	CustomerId	        string         `gorm:"type:varchar(64);unique;primaryKey" json:"customer_id" form:"customer_id"`
	CustomerName		string         `gorm:"type:varchar(80);not null" json:"customer_name" form:"customer_name"`
	Email 				string         `gorm:"type:varchar(50);unique;not null" json:"email" form:"email"`
	PhoneNumber 		string         `gorm:"type:varchar(20);unique;not null" json:"phone_number" form:"phone_number"`
	Dob 				time.Time      `gorm:"not null" json:"dob" form:"dob"`
	Sex 				bool           `gorm:"not null" json:"sex" form:"sex"`
	Salt 				string         `gorm:"type:varchar(80);unique;not null" json:"salt" form:"salt"`
	Password 			string         `gorm:"type:varchar(400);not null" json:"password" form:"password"`
	CreatedDate   		time.Time      `gorm:"not null" json:"created_date" form:"created_date"`
	Token	 			string		   `json:"token" form:"token"`
	
}
