package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/servlicense/servlicense/api/auth"
	"github.com/servlicense/servlicense/api/licenses"
	"github.com/servlicense/servlicense/api/types"
)

func CheckLicense(c *fiber.Ctx) error {
	// get param from url
	license := c.Params("license")

	// check if license is valid
	valid, err := licenses.CheckLicense(license)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(types.ApiResponse{
			Success: false,
			Code:    fiber.StatusInternalServerError,
			Message: err.Error(),
			Data:    nil,
		})
	}

	return c.Status(fiber.StatusOK).JSON(types.ApiResponse{
		Success: true,
		Code:    fiber.StatusOK,
		Message: "License checked successfully",
		Data: fiber.Map{
			"valid": valid,
		},
	})
}

func ListLicenses(c *fiber.Ctx) error {
	scopes := c.Locals("scopes").([]string)

	if scopes == nil {
		return c.JSON(types.ApiResponse{
			Success: false,
			Code:    401,
			Message: "Unauthorized, invalid api key",
			Data:    nil,
		})
	}

	// scopes should include admin or list_licenses, use auth.HasScope
	if !auth.HasScope(c, "admin") && !auth.HasScope(c, "list_licenses") {
		return c.JSON(types.ApiResponse{
			Success: false,
			Code:    403,
			Message: "Forbidden, missing required scope",
			Data:    nil,
		})
	}

	appID := c.Params("app_id")

	if appID == "" {
		return c.Status(fiber.StatusBadRequest).JSON(types.ApiResponse{
			Success: false,
			Code:    fiber.StatusBadRequest,
			Message: "Missing app_id",
			Data:    nil,
		})
	}

	licenses, err := licenses.ListLicenses(appID)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(types.ApiResponse{
			Success: false,
			Code:    fiber.StatusInternalServerError,
			Message: err.Error(),
			Data:    nil,
		})
	}

	return c.Status(fiber.StatusOK).JSON(types.ApiResponse{
		Success: true,
		Code:    fiber.StatusOK,
		Message: "Licenses listed successfully",
		Data: fiber.Map{
			"licenses": licenses,
		},
	})
}
