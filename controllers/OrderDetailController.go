package controllers

import (
	"net/http"

	"retailStore/lib/db"
	"retailStore/models"

	"github.com/labstack/echo"
)

func GetOrderDetailsController(c echo.Context) error {
	orderDetails, e := db.GetOrderDetails()
	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    200,
		"status":  "success",
		"message": "success get orderDetails",
		"orderDetail": orderDetails,
	})
}

func GetOrderDetailByIdController(c echo.Context) error {
	id := c.Param("id")
	orderDetail, getErr := db.GetOrderDetailById(id)
	if getErr != nil {
		return echo.NewHTTPError(http.StatusBadRequest, getErr.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    200,
		"status":  "success",
		"message": "success get orderDetail by id",
		"orderDetail": orderDetail,
	})
}

func CreateOrderDetailController(c echo.Context) error {
	orderDetail := models.OrderDetail{}
	c.Bind(&orderDetail)
	orderDetails, e := db.CreateOrderDetail(&orderDetail)
	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    200,
		"status":  "success",
		"message": "success creating new orderDetail",
		"orderDetail": orderDetails,
	})
}

func UpdateOrderDetailByIdController(c echo.Context) error {
	id := c.Param("id")
	var UpdateOrderDetail models.OrderDetail
	c.Bind(&UpdateOrderDetail)
	orderDetail, updateErr := db.UpdateOrderDetail(id, &UpdateOrderDetail)
	if updateErr != nil {
		return echo.NewHTTPError(http.StatusBadRequest, updateErr.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    200,
		"status":  "success",
		"message": "success updating existing orderDetail by id",
		"orderDetail": orderDetail,
	})
}

func DeleteOrderDetailByIdController(c echo.Context) error {
	id := c.Param("id")
	if _, deleteErr := db.DeleteOrderDetail(id); deleteErr != nil {
		return echo.NewHTTPError(http.StatusBadRequest, deleteErr.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    200,
		"status":  "success",
		"message": "success deleting existing orderDetail by id",
	})
}

