package main

import (
	"github.com/arossmann/24h-regional-api/store"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"log"
	"os"
)

func HealthGet(c *fiber.Ctx) error {
	return c.SendString("status: UP")
}
func baseRoute(c *fiber.Ctx) error {
	return c.SendString("API can be found at /api/v1")
}

func setupRoutes(app *fiber.App) {
	app.Get("/", baseRoute)
	app.Get("/health", HealthGet)
	app.Get("/api/v1/stores", store.GetStores)
	app.Get("/api/v1/stores/:id", store.GetStore)
	app.Post("/api/v1/stores", store.NewStore)
	app.Delete("/api/v1/stores/:id", store.DeleteStore)
}

func main() {
	app := fiber.New()
	// Default config
	app.Use(cors.New())

	// Or extend your config for customization
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))
	setupRoutes(app)
	log.Fatal(app.Listen(":" + os.Getenv("PORT")))

}
