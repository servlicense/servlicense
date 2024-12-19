package handlers

import (
	"github.com/Intevel/servlicense.sh/auth"
	"github.com/Intevel/servlicense.sh/database"
	"github.com/Intevel/servlicense.sh/types"
	"github.com/gofiber/fiber/v2"
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
			"api_keys": apiKeys,
		},
	})

}
