package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"retailStore/config"
	"retailStore/lib/db"
	"retailStore/models"
	"testing"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

func InitEchos() *echo.Echo {
	config.InitDBTest()
	e := echo.New()

	return e
}

func insertDB() {
	db.RegisterCustomer(&models.Customer{CustomerName: "sandi", Email: "sandipermana88@gmail.com"})

}

func TestGetCustomerController(t *testing.T) {
	var testCases = []struct {
		url         string
		code        int
		status      string
		expectedLen int
		isParam     bool
	}{
		{
			"/register",
			200,
			"success",
			3,
			false,
		},
		{
			"/register",
			200,
			"success",
			2,
			true,
		},
	}
	e := InitEcho()
	insertDB()
	for _, testCase := range testCases {
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		if testCase.isParam {
			q := make(url.Values)
			q.Set("category", "1")
			req = httptest.NewRequest(http.MethodGet, "/?"+q.Encode(), nil)
		}
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath(testCase.url)
		if assert.NoError(t, GetProductsController(c)) {
			assert.Equal(t, testCase.code, rec.Code)
			body := rec.Body.String()
			jsonData := []byte(body)
			var response models.Customer
			json.Unmarshal(jsonData, &response)
			fmt.Println(response)
			assert.Equal(t, testCase.expectedLen, len(response.CustomerId), "Product not selected based on category")
		}
	}
}

func TestErrorGetProductController(t *testing.T) {
	var testCases = []struct {
		url        string
		code       int
		status     string
		paramName  string
		paramValue string
	}{
		{
			"/products",
			400,
			"fail",
			"category",
			"nina",
		},
		{
			"/products",
			400,
			"fail",
			"category",
			"zzz",
		},
	}

	e := InitEcho()
	// insertDB()
	for _, testCase := range testCases {
		fmt.Println(testCase)
		q := make(url.Values)
		q.Set(testCase.paramName, testCase.paramValue)
		req := httptest.NewRequest(http.MethodGet, "/?"+q.Encode(), nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath(testCase.url)
		fmt.Println(q.Encode())
		assert.Error(t, GetProductsController(c))
		// assert.Equal(t,testCase.code,rec.Code)
		// body := rec.Body.String()
		// jsonData := []byte(body)
		// // var response models.ErrorResponse
		// var response models.ErrorResponse
		// json.Unmarshal(jsonData,&response)
		// fmt.Println(body)
		// assert.Equal(t,testCase.code,response.Code,"Error code must be 400")
		// assert.Equal(t,testCase.status,response.Status,"Error status must be fail")

	}

}
