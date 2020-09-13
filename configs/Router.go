package configs

import (
	"sync"

	"github.com/gofiber/fiber"
)

// IAppRouter is an interface for AppRouter
type IAppRouter interface {
	InitAppRouter(app *fiber.App) *fiber.App
}

type router struct{}

// InitRouter for route initialization
func (router *router) InitAppRouter(app *fiber.App) *fiber.App {
	api := app.Group("/api")
	api.Get("/healthcheck", func(c *fiber.Ctx) {
		c.Status(200).Send("")
	})

	positionsController := ServiceContainer().InjectPositionsController()
	positionsAPI := api.Group("/positions")
	positionsAPI.Get("/", positionsController.All)
	positionsAPI.Get("/:id", positionsController.Show)
	// positionsAPI.Post("/", positionsController.Create)
	// positionsAPI.Patch("/:id", positionsController.Update)
	// positionsAPI.Delete("/:id", positionsController.Delete)

	return app
}

var (
	r          *router
	routerOnce sync.Once
)

// AppRouter is a function to build routes
func AppRouter() IAppRouter {
	if r == nil {
		routerOnce.Do(func() {
			r = &router{}
		})
	}
	return r
}
