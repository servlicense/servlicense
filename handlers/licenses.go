package handlers

import (
	"github.com/Intevel/servlicense.sh/licenses"
	"github.com/Intevel/servlicense.sh/types"
	"github.com/gofiber/fiber/v2"
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
