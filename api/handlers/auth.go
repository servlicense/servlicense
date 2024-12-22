package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/servlicense/servlicense/api/auth"
	"github.com/servlicense/servlicense/api/database"
	"github.com/servlicense/servlicense/api/types"
)

func GetApiKeyInfo(c *fiber.Ctx) error {
	scopes := c.Locals("scopes")

	if scopes == nil {
		return c.JSON(types.ApiResponse{
			Success: false,
			Code:    401,
			Message: "Unauthorized, invalid api key",
			Data:    nil,
		})
	}

	return c.JSON(types.ApiResponse{
		Success: true,
		Code:    200,
		Message: "Successfully authenticated",
		Data: fiber.Map{
			"scopes": scopes,
		},
	})

}

func ListApiKeys(c *fiber.Ctx) error {
	scopes := c.Locals("scopes").([]string)

	if scopes == nil {
		return c.JSON(types.ApiResponse{
			Success: false,
			Code:    401,
			Message: "Unauthorized, invalid api key",
			Data:    nil,
		})
	}

	// scopes should include admin or manage_api_keys, use auth.HasScope
	if !auth.HasScope(c, "admin") && !auth.HasScope(c, "manage_api_keys") {
		return c.JSON(types.ApiResponse{
			Success: false,
			Code:    403,
			Message: "Forbidden, missing required scope",
			Data:    nil,
		})
	}

	apiKeys, err := database.Get().ListApiKeys()
	if err != nil {
		return c.JSON(types.ApiResponse{
			Success: false,
			Code:    500,
			Message: "Failed to list api keys",
			Data:    nil,
		})
	}

	return c.JSON(types.ApiResponse{
		Success: true,
		Code:    200,
		Message: "Successfully listed api keys",
		Data: fiber.Map{
			"apiKeys": apiKeys,
		},
	})

}

func CreateApiKey(c *fiber.Ctx) error {
	scopes := c.Locals("scopes").([]string)

	if scopes == nil {
		return c.JSON(types.ApiResponse{
			Success: false,
			Code:    401,
			Message: "Unauthorized, invalid api key",
			Data:    nil,
		})
	}
	var requestBody struct {
		Name   string   `json:"name"`
		Scopes []string `json:"scopes"`
	}

	if err := c.BodyParser(&requestBody); err != nil {
		return c.JSON(types.ApiResponse{
			Success: false,
			Code:    400,
			Message: "Invalid request body",
			Data:    nil,
		})
	}
	// scopes should include admin or manage_api_keys, use auth.HasScope
	if !auth.HasScope(c, "admin") && !auth.HasScope(c, "manage_api_keys") {
		return c.JSON(types.ApiResponse{
			Success: false,
			Code:    403,
			Message: "Forbidden, missing required scope",
			Data:    nil,
		})
	}

	identifier, apiKey, err := auth.CreateApiKey(requestBody.Name, requestBody.Scopes)
	if err != nil {
		return c.JSON(types.ApiResponse{
			Success: false,
			Code:    500,
			Message: "Failed to create api key",
			Data:    nil,
		})
	}

	return c.JSON(types.ApiResponse{
		Success: true,
		Code:    200,
		Message: "Successfully created api key",
		Data: fiber.Map{
			"id":      identifier,
			"api_key": apiKey,
		},
	})
}
