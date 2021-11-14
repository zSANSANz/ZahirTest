package middlewares

import (
	"os"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

func CreateToken(customerId string) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["customerId"] = customerId
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("SECRET_JWT")))
}

func ExtractTokenCustomerId(c echo.Context) float64 {
	customer := c.Get("customer").(*jwt.Token)
	if customer.Valid {
		claims := customer.Claims.(jwt.MapClaims)
		customerId := claims["customerId"].(float64)
		return customerId
	}
	return 0
}

func ExtractTokenCustomerRole(c echo.Context) string {
	customer := c.Get("customer").(*jwt.Token)
	if customer.Valid {
		claims := customer.Claims.(jwt.MapClaims)
		customerRole := claims["Role"].(string)
		return customerRole
	}
	return ""
}