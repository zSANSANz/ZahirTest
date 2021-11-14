package db

import (
	"retailStore/config"
	"retailStore/models"
)

func GetPaymentMethods() (interface{}, error) {
	var paymentMethod []models.PaymentMethod

	if err := config.DB.Find(&paymentMethod).Error; err != nil {
		return nil, err
	}
	return paymentMethod, nil
}

func GetPaymentMethodById(id string) (interface{}, error) {
	var paymentMethod []models.PaymentMethod

	if err := config.DB.First(&paymentMethod, "payment_method_id=?", id).Error; err != nil {
		return nil, err
	}
	return paymentMethod, nil
}

func CreatePaymentMethod(paymentMethod *models.PaymentMethod) (interface{}, error) {
	if err := config.DB.Save(paymentMethod).Error; err != nil {
		return nil, err
	}
	return paymentMethod, nil
}

func UpdatePaymentMethod(id string, paymentMethod *models.PaymentMethod) (interface{}, error) {
	var existingPaymentMethod models.PaymentMethod
	if err := config.DB.First(&existingPaymentMethod, "payment_method_id=?", id).Error; err != nil {
		return nil, err
	}
	existingPaymentMethod.MethodName = paymentMethod.MethodName
	existingPaymentMethod.Code = paymentMethod.Code
	if updateErr := config.DB.Save(&existingPaymentMethod).Where("payment_method_id=?", id).Error; updateErr != nil {
		return nil, updateErr
	}
	return existingPaymentMethod, nil
}

func DeletePaymentMethod(id string) (interface{}, error) {
	var paymentMethod models.PaymentMethod
	if err := config.DB.First(&paymentMethod, "payment_method_id=?", id).Error; err != nil {
		return nil, err
	}
	if deleteErr := config.DB.Delete(&paymentMethod).Where("payment_method_id=?", id).Error; deleteErr != nil {
		return nil, deleteErr
	}
	return paymentMethod, nil
}
