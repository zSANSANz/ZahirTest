package db

import (
	"retailStore/config"
	"retailStore/models"
	"time"
)

func GetOrders() (interface{}, error) {
	var order []models.Order

	if err := config.DB.Find(&order).Error; err != nil {
		return nil, err
	}
	return order, nil
}

func GetOrderById(id string) (interface{}, error) {
	var order []models.Order

	if err := config.DB.First(&order, "order_id=?", id).Error; err != nil {
		return nil, err
	}
	return order, nil
}

func CreateOrder(order *models.Order) (interface{}, error) {
	var orders models.Order

	currentTime := time.Now()

	thisMonth := ""

	switch currentTime.Format("1") {
    case "1":
        thisMonth = "I"
    case "2":
        thisMonth = "II"
    case "3":
		thisMonth = "III"
	case "4":
        thisMonth = "IV"
    case "5":
		thisMonth = "V"
	case "6":
        thisMonth = "VI"
    case "7":
		thisMonth = "VII"
    case "8":
		thisMonth = "VIII"
	case "9":
        thisMonth = "IX"
    case "10":
		thisMonth = "X"
	case "11":
        thisMonth = "XI"
    case "12":
		thisMonth = "XII"
    }

    orders.OrderId = "PO-123/" + thisMonth + "/" + currentTime.Format("2006")
	orders.CustomerId = order.CustomerId
	orders.OrderNumber = order.OrderNumber
	orders.OrderDate = time.Now()
	orders.PaymentMethod = order.PaymentMethod

	if err := config.DB.Save(orders).Error; err != nil {
		return nil, err
	}
	return orders, nil
}

func UpdateOrder(id string, order *models.Order) (interface{}, error) {
	var existingOrder models.Order
	if err := config.DB.First(&existingOrder, "order_id=?", id).Error; err != nil {
		return nil, err
	}
	existingOrder.CustomerId = order.CustomerId
	existingOrder.OrderNumber = order.OrderNumber
	if updateErr := config.DB.Save(&existingOrder).Where("order_id=?", id).Error; updateErr != nil {
		return nil, updateErr
	}
	return existingOrder, nil
}

func DeleteOrder(id string) (interface{}, error) {
	var order models.Order
	if err := config.DB.First(&order, "order_id=?", id).Error; err != nil {
		return nil, err
	}
	if deleteErr := config.DB.Delete(&order).Where("order_id=?", id).Error; deleteErr != nil {
		return nil, deleteErr
	}
	return order, nil
}
