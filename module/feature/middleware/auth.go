package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	users "ruti-store/module/feature/user/domain"
	"ruti-store/utils/response"
	"ruti-store/utils/token"

	"strings"
)

// AuthMiddleware is a middleware for JWT authentication
func AuthMiddleware(jwtService token.JWTInterface, userService users.UserServiceInterface) fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")

		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			return response.ErrorBuildResponse(c, fiber.StatusForbidden, "Forbidden: Missing or invalid authorization header.")
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		tokens, err := jwtService.ValidateToken(tokenString)
		if err != nil {
			return response.ErrorBuildResponse(c, fiber.StatusForbidden, "Forbidden: Invalid token "+err.Error())
		}

		claims, ok := tokens.Claims.(jwt.MapClaims)
		if !ok || !tokens.Valid {
			return response.ErrorBuildResponse(c, fiber.StatusForbidden, "Forbidden: Invalid or expired token.")
		}

		userIDFloat, ok := claims["user_id"].(float64)
		if !ok {
			return response.ErrorBuildResponse(c, fiber.StatusForbidden, "Forbidden: Invalid user ID in token.")
		}

		userID := uint64(userIDFloat)

		user, err := userService.GetUserByID(userID)
		if err != nil {
			return response.ErrorBuildResponse(c, fiber.StatusForbidden, "Forbidden: User not found.")
		}

		c.Locals("currentUser", user)

		return c.Next()
	}
}
