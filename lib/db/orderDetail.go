package db

import (
	"retailStore/config"
	"retailStore/models"
)

func GetOrderDetails() (interface{}, error) {
	var orderDetail []models.OrderDetail

	if err := config.DB.Find(&orderDetail).Error; err != nil {
		return nil, err
	}
	return orderDetail, nil
}

func GetOrderDetailById(id string) (interface{}, error) {
	var orderDetail []models.OrderDetail

	if err := config.DB.First(&orderDetail, "order_detail_id=?", id).Error; err != nil {
		return nil, err
	}
	return orderDetail, nil
}

func CreateOrderDetail(orderDetail *models.OrderDetail) (interface{}, error) {
	if err := config.DB.Save(orderDetail).Error; err != nil {
		return nil, err
	}
	return orderDetail, nil
}

func UpdateOrderDetail(id string, orderDetail *models.OrderDetail) (interface{}, error) {
	var existingOrderDetail models.OrderDetail
	if err := config.DB.First(&existingOrderDetail, "order_detail_id=?", id).Error; err != nil {
		return nil, err
	}
	existingOrderDetail.OrderId = orderDetail.OrderId
	existingOrderDetail.ProductId = orderDetail.ProductId
	existingOrderDetail.Qty = orderDetail.Qty
	if updateErr := config.DB.Save(&existingOrderDetail).Where("order_detail_id=?", id).Error; updateErr != nil {
		return nil, updateErr
	}
	return existingOrderDetail, nil
}

func DeleteOrderDetail(id string) (interface{}, error) {
	var orderDetail models.OrderDetail
	if err := config.DB.First(&orderDetail, "order_detail_id=?", id).Error; err != nil {
		return nil, err
	}
	if deleteErr := config.DB.Delete(&orderDetail).Where("order_detail_id=?", id).Error; deleteErr != nil {
		return nil, deleteErr
	}
	return orderDetail, nil
}
