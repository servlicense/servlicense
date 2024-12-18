package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"

	"github.com/Intevel/servlicense.sh/config"
	"github.com/Intevel/servlicense.sh/database"
	"github.com/Intevel/servlicense.sh/types"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	err := config.LoadConfig("config.toml")
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	cfg := config.GetConfig()

	log.Printf("Loaded config: %v", cfg)

	db := database.Get()
	err = db.Connect(filepath.Join("servlicense.db"))

	if err != nil {
		fmt.Println("Failed to connect to database:", err)
		os.Exit(1)
	}
	if err := db.CreateTablesIfNotExist(); err != nil {
		fmt.Println("Failed to create tables:", err)
	}

	app := fiber.New(fiber.Config{
		AppName:           "servlicense.sh",
		ServerHeader:      "servlicense.sh",
		EnablePrintRoutes: false,
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			code := fiber.StatusInternalServerError

			var e *fiber.Error
			if errors.As(err, &e) {
				code = e.Code
			}

			return c.Status(code).JSON(types.ApiResponse{
				Success: false,
				Code:    code,
				Message: err.Error(),
				Data:    nil,
			})
		},
	})
	app.Use(cors.New(cors.Config{
		AllowMethods:     "",
		Next:             nil,
		AllowOrigins:     "*",
		AllowHeaders:     "",
		AllowCredentials: false,
		ExposeHeaders:    "",
		MaxAge:           0,
	}))

	app.Use(logger.New(logger.Config{
		TimeFormat: "2006-01-02 15:04:05",
	}))

	app.Use(func(c *fiber.Ctx) error {
		return c.Status(404).JSON(types.ApiResponse{
			Code:    fiber.ErrNotFound.Code,
			Message: fiber.ErrNotFound.Message,
			Success: false,
		})
	})

	log.Fatal(app.Listen(
		cfg.Host + ":" + strconv.Itoa(cfg.Port),
	))
}