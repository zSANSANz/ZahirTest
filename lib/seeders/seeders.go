package seeders

import (
	"retailStore/config"
	"retailStore/models"
)

func Seed() error {
	if _, err2 := ProductsSeed(); err2 != nil {
		return err2
	}
	return nil
}

func ProductsSeed() ([]models.Product, error) {
	ProductNames := []string{"JNE", "J&T", "TIKI", "JET", "Wahana"}
	for _, productName := range ProductNames {
		product := models.Product{
			ProductName: productName,
		}
		err := config.DB.Create(&product).Error
		if err != nil {
			return nil, err
		}
	}
	products := []models.Product{}
	//config.DB.Find(&couriers)

	return products, nil
}
