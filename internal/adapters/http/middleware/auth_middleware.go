package middleware

import (
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/spf13/viper"
)

func CreateToken(userID int, email string, role string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = userID
	claims["email"] = email
	claims["role"] = role
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	secret := viper.GetString("SECRET_JWT")
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
func RequireRole(role string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			user := c.Get("user").(*jwt.Token)
			claims := user.Claims.(*jwt.MapClaims)
			userRole := (*claims)["role"].(string)

			if userRole != role {
				return c.JSON(http.StatusUnauthorized, map[string]string{"message": "Unauthorized access"})
			}

			return next(c)
		}
	}
}

func AuthMiddleware() echo.MiddlewareFunc {
	return middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey: []byte(viper.GetString("SECRET_JWT")),
		ErrorHandler: func(err error) error {
			if _, ok := err.(*jwt.ValidationError); ok {
				return echo.NewHTTPError(http.StatusUnauthorized, "Invalid token")
			}
			return err
		},
		Claims: &jwt.MapClaims{},
	})
}
