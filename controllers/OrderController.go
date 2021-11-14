package controllers

import (
	"net/http"

	"retailStore/lib/db"
	"retailStore/models"

	"github.com/labstack/echo"
)

func GetOrdersController(c echo.Context) error {
	orders, e := db.GetOrders()
	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    200,
		"status":  "success",
		"message": "success get orders",
		"order": orders,
	})
}

func GetOrderByIdController(c echo.Context) error {
	id := c.Param("id")
	order, getErr := db.GetOrderById(id)
	if getErr != nil {
		return echo.NewHTTPError(http.StatusBadRequest, getErr.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    200,
		"status":  "success",
		"message": "success get order by id",
		"order": order,
	})
}

func CreateOrderController(c echo.Context) error {
	order := models.Order{}
	c.Bind(&order)
	orders, e := db.CreateOrder(&order)
	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    200,
		"status":  "success",
		"message": "success creating new order",
		"order": orders,
	})
}

func UpdateOrderByIdController(c echo.Context) error {
	id := c.Param("id")
	var UpdateOrder models.Order
	c.Bind(&UpdateOrder)
	order, updateErr := db.UpdateOrder(id, &UpdateOrder)
	if updateErr != nil {
		return echo.NewHTTPError(http.StatusBadRequest, updateErr.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    200,
		"status":  "success",
		"message": "success updating existing order by id",
		"order": order,
	})
}

func DeleteOrderByIdController(c echo.Context) error {
	id := c.Param("id")
	if _, deleteErr := db.DeleteOrder(id); deleteErr != nil {
		return echo.NewHTTPError(http.StatusBadRequest, deleteErr.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    200,
		"status":  "success",
		"message": "success deleting existing order by id",
	})
}

