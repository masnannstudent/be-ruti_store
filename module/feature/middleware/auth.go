package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"ruti-store/utils/response"
	"ruti-store/utils/token"
	"strings"
)

func Auth(jwtService token.JWTInterface) fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")
		if authHeader == "" || len(strings.Split(authHeader, "Bearer ")) != 2 {
			return response.ErrorBuildResponse(c, fiber.StatusUnauthorized, "Unauthorized: Invalid or missing authorization header.")
		}

		tokenString := strings.Split(authHeader, "Bearer ")[1]

		validToken, err := jwtService.ValidateToken(tokenString)
		if err != nil || !validToken.Valid {
			return response.ErrorBuildResponse(c, fiber.StatusUnauthorized, "Unauthorized: Invalid token")
		}

		claims, ok := validToken.Claims.(jwt.MapClaims)
		if !ok {
			return response.ErrorBuildResponse(c, fiber.StatusUnauthorized, "Unauthorized: Unable to extract user information from token.")
		}

		currentUser := claims["user"].(string)

		// Menyimpan informasi pengguna dalam context Fiber
		c.Locals("currentUser", currentUser)

		return c.Next()
	}
}
