package routes

import (
	"fiber/controllers"
	"fiber/middlewares"
	"github.com/gofiber/fiber/v2"
)

func RegisterAPI(app *fiber.App) {
	registerIndex(app)
	registerUsers(app)
	registerTests(app)
}

func registerIndex(app *fiber.App) {

	app.Get("/", func(c *fiber.Ctx) error {
		// Render index template
		return c.Render("index", fiber.Map{
			"Title": "Hello, World!",
		})
	})
}

func registerUsers(app *fiber.App) {
	v1 := app.Group("/v1")

	usersV1 := v1.Group("/users")
	usersV1.Use(middlewares.UserMiddleware) // add middleware
	usersV1.Get("/get-one", controllers.GetOne)
}

func registerTests(app *fiber.App) {
	t := app.Group("/test")
	t.Get("/test", controllers.Test)
	t.Get("/test-orm", controllers.TestOrm)
	t.Get("/test-redis", controllers.TestRedis)
	t.Get("/test-fast-http", controllers.TestFastHttp)
	t.Get("/test-mongo-db", controllers.TestMongoDb)
}