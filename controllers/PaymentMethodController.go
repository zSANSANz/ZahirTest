package controllers

import (
	"net/http"

	"retailStore/lib/db"
	"retailStore/models"

	"github.com/labstack/echo"
)

func GetPaymentMethodsController(c echo.Context) error {
	paymentMethods, e := db.GetPaymentMethods()
	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    200,
		"status":  "success",
		"message": "success get paymentMethods",
		"paymentMethod": paymentMethods,
	})
}

func GetPaymentMethodByIdController(c echo.Context) error {
	id := c.Param("id")
	paymentMethod, getErr := db.GetPaymentMethodById(id)
	if getErr != nil {
		return echo.NewHTTPError(http.StatusBadRequest, getErr.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    200,
		"status":  "success",
		"message": "success get paymentMethod by id",
		"paymentMethod": paymentMethod,
	})
}

func CreatePaymentMethodController(c echo.Context) error {
	paymentMethod := models.PaymentMethod{}
	c.Bind(&paymentMethod)
	paymentMethods, e := db.CreatePaymentMethod(&paymentMethod)
	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    200,
		"status":  "success",
		"message": "success creating new paymentMethod",
		"paymentMethod": paymentMethods,
	})
}

func UpdatePaymentMethodByIdController(c echo.Context) error {
	id := c.Param("id")
	var UpdatePaymentMethod models.PaymentMethod
	c.Bind(&UpdatePaymentMethod)
	paymentMethod, updateErr := db.UpdatePaymentMethod(id, &UpdatePaymentMethod)
	if updateErr != nil {
		return echo.NewHTTPError(http.StatusBadRequest, updateErr.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    200,
		"status":  "success",
		"message": "success updating existing paymentMethod by id",
		"paymentMethod": paymentMethod,
	})
}

func DeletePaymentMethodByIdController(c echo.Context) error {
	id := c.Param("id")
	if _, deleteErr := db.DeletePaymentMethod(id); deleteErr != nil {
		return echo.NewHTTPError(http.StatusBadRequest, deleteErr.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    200,
		"status":  "success",
		"message": "success deleting existing paymentMethod by id",
	})
}

