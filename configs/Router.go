package configs

import (
	"sync"

	"github.com/gofiber/fiber/v2"
)

// IAppRouter is an interface for AppRouter
type IAppRouter interface {
	RegisterAPI(app *fiber.App)
}

type router struct{}

// RegisterAPI is a function to initialize API routes
func (router *router) RegisterAPI(app *fiber.App) {
	api := app.Group("/api")
	api.Get("/healthcheck", func(c *fiber.Ctx) error {
		return c.Status(200).SendString("")
	})

	apiV1 := api.Group("/v1")
	RegisterPositionsAPI(apiV1)
	RegisterClubsAPI(apiV1)
}

// RegisterPositionsAPI is a function to initialize Positions API
func RegisterPositionsAPI(api fiber.Router) {
	positionsController := ServiceContainer().InjectPositionsController()
	positionsAPI := api.Group("/positions")
	positionsAPI.Get("/", positionsController.All)
	positionsAPI.Get("/:id", positionsController.Show)
	positionsAPI.Post("/", positionsController.Create)
	positionsAPI.Post("/create_many", positionsController.CreateMany)
	positionsAPI.Patch("/:id", positionsController.Update)
	positionsAPI.Delete("/:id", positionsController.Delete)
}

// RegisterClubsAPI is a function to initialize Clubs API
func RegisterClubsAPI(api fiber.Router) {
	clubsController := ServiceContainer().InjectClubsController()
	clubsAPI := api.Group("/clubs")
	clubsAPI.Get("/", clubsController.All)
	clubsAPI.Get("/:id", clubsController.Show)
	clubsAPI.Post("/", clubsController.Create)
	clubsAPI.Post("/create_many", clubsController.CreateMany)
	clubsAPI.Patch("/:id", clubsController.Update)
	clubsAPI.Delete("/:id", clubsController.Delete)
}

var (
	r          *router
	routerOnce sync.Once
)

// AppRouter is a function to initialize IAppRouter instance
func AppRouter() IAppRouter {
	if r == nil {
		routerOnce.Do(func() {
			r = &router{}
		})
	}
	return r
}
