package controllers

import (
	"net/http"

	"retailStore/lib/db"
	"retailStore/models"

	"github.com/labstack/echo"
)

func GetCustomersController(c echo.Context) error {
	customers, e := db.GetCustomers()
	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    200,
		"status":  "success",
		"message": "success get customers",
		"customer": customers,
	})
}

func GetCustomerByIdController(c echo.Context) error {
	id := c.Param("id")
	customer, getErr := db.GetCustomerById(id)
	if getErr != nil {
		return echo.NewHTTPError(http.StatusBadRequest, getErr.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    200,
		"status":  "success",
		"message": "success get customer by id",
		"customer": customer,
	})
}

func CreateCustomerController(c echo.Context) error {
	customer := models.Customer{}
	c.Bind(&customer)
	customers, e := db.CreateCustomer(&customer)
	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    200,
		"status":  "success",
		"message": "success creating new customer",
		"customer": customers,
	})
}

func UpdateCustomerByIdController(c echo.Context) error {
	id := c.Param("id")
	var UpdateCustomer models.Customer
	c.Bind(&UpdateCustomer)
	customer, updateErr := db.UpdateCustomer(id, &UpdateCustomer)
	if updateErr != nil {
		return echo.NewHTTPError(http.StatusBadRequest, updateErr.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    200,
		"status":  "success",
		"message": "success updating existing customer by id",
		"customer": customer,
	})
}

func DeleteCustomerByIdController(c echo.Context) error {
	id := c.Param("id")
	if _, deleteErr := db.DeleteCustomer(id); deleteErr != nil {
		return echo.NewHTTPError(http.StatusBadRequest, deleteErr.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    200,
		"status":  "success",
		"message": "success deleting existing customer by id",
	})
}

func RegisterCustomerController(c echo.Context) error {
	customer := models.Customer{}
	c.Bind(&customer)

	_, err := db.RegisterCustomer(&customer)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"customer":   customer,
	})
}

func LoginCustomerController(c echo.Context) error {
	customer := models.Customer{}
	c.Bind(&customer)

	customers, e := db.LoginCustomer(&customer)
	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success login",
		"customers": customers,
	})
}

