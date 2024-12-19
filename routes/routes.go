package routes

import (
	"log"

	"github.com/Intevel/servlicense.sh/auth"
	"github.com/Intevel/servlicense.sh/handlers"
	"github.com/gofiber/fiber/v2"
)

type Route struct {
	Path    string
	Method  string
	Handler func(*fiber.Ctx) error
}

// Unauthenticated routes (no JWT check)
var UnauthenticatedRoutes = []Route{}

// Authenticated routes (api key check required)
var AuthenticatedRoutes = []Route{
	{
		Path:    "/ping",
		Method:  "GET",
		Handler: handlers.Ping,
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
	group := app.Group(groupPrefix, auth.AuthMiddleware()) // Authenticated group with JWT middleware
	for _, route := range routes {
		group.Add(route.Method, route.Path, route.Handler)
		log.Printf("Registered authenticated route: [%s] %s%s\n", route.Method, groupPrefix, route.Path)
	}
	log.Printf("Registered '%d' authenticated routes\n", len(routes))
}
