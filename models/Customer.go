package models

type Customer struct {
	CustomerId   string `gorm:"type:varchar(64);unique;primaryKey" json:"customer_id" form:"customer_id"`
	CustomerName string `gorm:"type:varchar(80);not null" json:"customer_name" form:"customer_name"`
	Email        string `gorm:"type:varchar(50);unique;not null" json:"email" form:"email"`
	PhoneNumber  string `gorm:"type:varchar(20);unique;not null" json:"phone_number" form:"phone_number"`
	Sex          bool   `gorm:"not null" json:"sex" form:"sex"`
	Salt         string `gorm:"type:varchar(80);unique;not null" json:"salt" form:"salt"`
	Password     string `gorm:"type:varchar(400);not null" json:"password" form:"password"`
	Token        string `json:"token" form:"token"`
}
