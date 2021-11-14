package controllers

import (
	"net/http"

	"retailStore/lib/db"
	"retailStore/models"

	"github.com/labstack/echo"
)

func GetProductsController(c echo.Context) error {
	products, e := db.GetProducts()
	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    200,
		"status":  "success",
		"message": "success get products",
		"product": products,
	})
}

func GetProductByIdController(c echo.Context) error {
	id := c.Param("id")
	product, getErr := db.GetProductById(id)
	if getErr != nil {
		return echo.NewHTTPError(http.StatusBadRequest, getErr.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    200,
		"status":  "success",
		"message": "success get product by id",
		"product": product,
	})
}

func CreateProductController(c echo.Context) error {
	product := models.Product{}
	c.Bind(&product)
	products, e := db.CreateProduct(&product)
	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    200,
		"status":  "success",
		"message": "success creating new product",
		"product": products,
	})
}

func UpdateProductByIdController(c echo.Context) error {
	id := c.Param("id")
	var UpdateProduct models.Product
	c.Bind(&UpdateProduct)
	product, updateErr := db.UpdateProduct(id, &UpdateProduct)
	if updateErr != nil {
		return echo.NewHTTPError(http.StatusBadRequest, updateErr.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    200,
		"status":  "success",
		"message": "success updating existing product by id",
		"product": product,
	})
}

func DeleteProductByIdController(c echo.Context) error {
	id := c.Param("id")
	if _, deleteErr := db.DeleteProduct(id); deleteErr != nil {
		return echo.NewHTTPError(http.StatusBadRequest, deleteErr.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    200,
		"status":  "success",
		"message": "success deleting existing product by id",
	})
}

