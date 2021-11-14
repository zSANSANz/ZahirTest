package db

import (
	"retailStore/config"
	"retailStore/models"
	"retailStore/middlewares"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"time"
)

func GetCustomers() (interface{}, error) {
	var customer []models.Customer

	if err := config.DB.Find(&customer).Error; err != nil {
		return nil, err
	}
	return customer, nil
}

func GetCustomerById(id string) (interface{}, error) {
	var customer []models.Customer

	if err := config.DB.First(&customer, "customer_id=?", id).Error; err != nil {
		return nil, err
	}
	return customer, nil
}

func CreateCustomer(customer *models.Customer) (interface{}, error) {
	if err := config.DB.Save(customer).Error; err != nil {
		return nil, err
	}
	return customer, nil
}

func UpdateCustomer(id string, customer *models.Customer) (interface{}, error) {
	var existingCustomer models.Customer
	if err := config.DB.First(&existingCustomer, "customer_id=?", id).Error; err != nil {
		return nil, err
	}
	existingCustomer.CustomerName = customer.CustomerName
	existingCustomer.Email = customer.Email
	existingCustomer.PhoneNumber = customer.PhoneNumber
	existingCustomer.Sex = customer.Sex
	existingCustomer.Salt = customer.Salt
	existingCustomer.Password = customer.Password
	if updateErr := config.DB.Save(&existingCustomer).Where("customer_id=?", id).Error; updateErr != nil {
		return nil, updateErr
	}
	return existingCustomer, nil
}

func DeleteCustomer(id string) (interface{}, error) {
	var customer models.Customer
	if err := config.DB.First(&customer, "customer_id=?", id).Error; err != nil {
		return nil, err
	}
	if deleteErr := config.DB.Delete(&customer).Where("customer_id=?", id).Error; deleteErr != nil {
		return nil, deleteErr
	}
	return customer, nil
}

func RegisterCustomer(customer *models.Customer) (interface{}, error) {
	var customers models.Customer
	bytes, _ := bcrypt.GenerateFromPassword([]byte(customer.Password), 10)
	fmt.Println(string(bytes))
	customers.CustomerId = customer.CustomerId
	customers.CustomerName = customer.CustomerName
	customers.Email = customer.Email
	customers.PhoneNumber = customer.PhoneNumber
	customers.Salt = customer.Salt
	customers.Password = string(bytes)
	customers.CreatedDate = time.Now()

	if err := config.DB.Create(customers).Error; err != nil {
		return nil, err
	}

	var err error
	customer.Token, err = middlewares.CreateToken(string(customer.CustomerId))
	if err != nil {
		return nil, err
	}

	if err := config.DB.Save(customers).Error; err != nil {
		return nil, err
	}

	return customers, nil
}

func LoginCustomer(customer *models.Customer) (interface{}, error) {
	var customers models.Customer

	if err := config.DB.First(&customers, "email=?", customer.Email).Error; err != nil {
		return nil, err
	}

	err := bcrypt.CompareHashAndPassword([]byte(customers.Password), []byte(customer.Password))

	if err != nil {
		return nil, err
	}

	customer.Token, err = middlewares.CreateToken(string(customer.CustomerId))
	if err != nil {
		return nil, err
	}

	return customer, nil
}