package routes

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/servlicense/servlicense/api/auth"
	"github.com/servlicense/servlicense/api/handlers"
)

type Route struct {
	Path    string
	Method  string
	Handler func(*fiber.Ctx) error
}

// Unauthenticated routes (no JWT check)
var UnauthenticatedRoutes = []Route{
	{
		Path:    "/licenses/check/:license",
		Method:  "GET",
		Handler: handlers.CheckLicense,
	},
}

// Authenticated routes (api key check required)
var AuthenticatedRoutes = []Route{
	{
		Path:    "/auth/me",
		Method:  "GET",
		Handler: handlers.GetApiKeyInfo,
	},
	// Needs admin or manage_api_keys scope
	{
		Path:    "/auth/apikeys",
		Method:  "GET",
		Handler: handlers.ListApiKeys,
	},
	{
		Path:    "/auth/apikeys",
		Method:  "POST",
		Handler: handlers.CreateApiKey,
	},
	// Needs admin or list_licenses scope
	{
		Path:    "/:app_id/licenses",
		Method:  "GET",
		Handler: handlers.ListLicenses,
	},
}

func RegisterRoutes(app *fiber.App, groupPrefix string, routes ...Route) {
	group := app.Group(groupPrefix) // Unauthenticated group (no middleware)
	for _, route := range routes {
		group.Add(route.Method, route.Path, route.Handler)
		log.Printf("Registered unauthenticated route: [%s] %s%s\n", route.Method, groupPrefix, route.Path)
	}
	log.Printf("Registered '%d' unauthenticated routes\n", len(routes))
}

func RegisterAuthenticatedRoutes(app *fiber.App, groupPrefix string, routes ...Route) {
	group := app.Group(groupPrefix) // Authenticated group with middleware
	for _, route := range routes {
		group.Add(route.Method, route.Path, auth.AuthMiddleware(), route.Handler)
		log.Printf("Registered authenticated route: [%s] %s%s\n", route.Method, groupPrefix, route.Path)
	}
	log.Printf("Registered '%d' authenticated routes\n", len(routes))
}
