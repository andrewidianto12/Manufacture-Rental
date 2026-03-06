package middleware

import (
	"net/http"
	"strings"

	"github.com/andrewidianto12/Manufacture-Rental/util"
	"github.com/labstack/echo/v4"
)

func JWTMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			authHeader := c.Request().Header.Get("Authorization")
			if strings.TrimSpace(authHeader) == "" {
				return c.JSON(http.StatusUnauthorized, map[string]any{"message": "missing authorization header"})
			}

			parts := strings.SplitN(authHeader, " ", 2)
			if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" || strings.TrimSpace(parts[1]) == "" {
				return c.JSON(http.StatusUnauthorized, map[string]any{"message": "invalid authorization header format"})
			}

			userID, err := util.ParseToken(parts[1])
			if err != nil {
				return c.JSON(http.StatusUnauthorized, map[string]any{"message": "invalid or expired token", "error": err.Error()})
			}

			c.Set("user_id", userID)
			return next(c)
		}
	}
}
