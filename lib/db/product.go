package db

import (
	"retailStore/config"
	"retailStore/models"
)

func GetProducts() (interface{}, error) {
	var product []models.Product

	if err := config.DB.Find(&product).Error; err != nil {
		return nil, err
	}
	return product, nil
}

func GetProductById(id string) (interface{}, error) {
	var product []models.Product

	if err := config.DB.First(&product, "product_id=?", id).Error; err != nil {
		return nil, err
	}
	return product, nil
}

func CreateProduct(product *models.Product) (interface{}, error) {
	if err := config.DB.Save(product).Error; err != nil {
		return nil, err
	}
	return product, nil
}

func UpdateProduct(id string, product *models.Product) (interface{}, error) {
	var existingProduct models.Product
	if err := config.DB.First(&existingProduct, "product_id=?", id).Error; err != nil {
		return nil, err
	}
	existingProduct.ProductName = product.ProductName
	existingProduct.BasicPrice = product.BasicPrice
	if updateErr := config.DB.Save(&existingProduct).Where("product_id=?", id).Error; updateErr != nil {
		return nil, updateErr
	}
	return existingProduct, nil
}

func DeleteProduct(id string) (interface{}, error) {
	var product models.Product
	if err := config.DB.First(&product, "product_id=?", id).Error; err != nil {
		return nil, err
	}
	if deleteErr := config.DB.Delete(&product).Where("product_id=?", id).Error; deleteErr != nil {
		return nil, deleteErr
	}
	return product, nil
}
