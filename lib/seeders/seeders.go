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
	ProductNames := []string{"tepung", "beras", "telur", "mie instan", "minyak goreng"}
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

func CustomersSeed() ([]models.Customer, error) {
	CustomerNames := []string{"sandi", "joko", "kelly", "budi", "rihana"}
	for _, customerName := range CustomerNames {
		customer := models.Customer{
			CustomerName: customerName,
		}
		err := config.DB.Create(&customer).Error
		if err != nil {
			return nil, err
		}
	}
	customer := []models.Customer{}
	//config.DB.Find(&couriers)

	return customer, nil

}
