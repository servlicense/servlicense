package auth

import (
	"encoding/base64"
	"fmt"
	"strings"

	"github.com/Intevel/servlicense.sh/database"
	"github.com/gofiber/fiber/v2"
)

// AuthMiddleware is a Fiber middleware that checks for a valid API key in the Authorization header
func AuthMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Authorization header is missing",
			})
		}

		decoded, err := base64.StdEncoding.DecodeString(authHeader)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Authorization header is not properly base64 encoded",
			})
		}

		authHeaderParts := strings.Split(string(decoded), ":")
		if len(authHeaderParts) != 2 {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Authorization header is invalid",
			})
		}

		id := authHeaderParts[0]
		fmt.Println("id: ", id)
		apiKey := authHeaderParts[1]

		apiKeyData, err := database.Get().GetApiKey(id)
		if err != nil {
			fmt.Println("Error getting API key data: ", err)
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Invalid id",
			})
		}

		// Verify the API key
		ok, err := VerifyApiKey(apiKey, apiKeyData.ApiKey)
		if err != nil {
			fmt.Println("Error verifying API key: ", err)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Internal server error",
			})
		}

		if !ok {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Invalid API key",
			})
		}

		// Store scopes in the request context for later use
		c.Locals("scopes", apiKeyData.Scopes)
		return c.Next()
	}
}
