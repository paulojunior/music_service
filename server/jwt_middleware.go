package server

import (
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
)

// JWTMiddleware is the middleware for validating JWT tokens
func JWTMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		tokenString := c.Request().Header.Get("Authorization")
		if tokenString == "" {
			return echo.NewHTTPError(http.StatusUnauthorized, "Token not provided")
		}

		// Get the secret key from Viper
		secretKey := []byte(viper.GetString("jwt_secret"))

		tokenString = strings.Replace(tokenString, "Bearer ", "", 1)

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// Use the secret key from Viper to verify the token's signature
			return secretKey, nil
		})

		if err != nil || !token.Valid {
			return echo.NewHTTPError(http.StatusUnauthorized, "Invalid token")
		}

		// Checking the validity of the expiration time
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			// Valid expiration time
			exp := claims["exp"].(float64)
			if int64(exp) < time.Now().Unix() {
				// Token has expired
				return echo.NewHTTPError(http.StatusUnauthorized, "Token has expired")
			}
		}

		// If the token is valid, continue to the next handler
		return next(c)
	}
}
