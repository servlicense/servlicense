package auth

// write a go fiber middleware that checks if the request has a valid api key, the authorization header will be a base64 encoded string "id:api_key", use database.GetApiKey(id) for that

import (
	"encoding/base64"
	"strings"

	"github.com/Intevel/servlicense.sh/database"
	"github.com/gofiber/fiber/v2"
)

func AuthMiddleware(c *fiber.Ctx) error {
	authHeader := c.Get("Authorization")
	if authHeader == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Authorization header is missing",
		})
	}

	authHeaderParts := strings.Split(authHeader, ":")
	if len(authHeaderParts) != 2 {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Authorization header is invalid",
		})
	}

	decoded, err := base64.StdEncoding.DecodeString(authHeader)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Authorization header is not properly base64 encoded",
		})
	}

	authHeaderParts = strings.Split(string(decoded), ":")
	if len(authHeaderParts) != 2 {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Authorization header is invalid",
		})
	}

	id := authHeaderParts[0]
	apiKey := authHeaderParts[1]

	apiKeyData, err := database.Get().GetApiKey(id)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid id",
		})
	}

	if apiKeyData.ApiKey != apiKey {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid api",
		})
	}

	c.Locals("scopes", apiKeyData.Scopes)
	return c.Next()
}
