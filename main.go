package main

// unfinished

import (
	"github.com/gofiber/contrib/fiberzap/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/vt-d/audite/routes"
	"go.uber.org/zap"
)

func main() {
	app := fiber.New()
	logger, _ := zap.NewProduction()

	log.SetLogger(fiberzap.NewLogger(fiberzap.LoggerConfig{
		ExtraKeys: []string{"request_id"},
		SetLogger: logger,
	}))

	app.Use(fiberzap.New(fiberzap.Config{
		Logger: logger,
	}))

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	routes.InitWebsocket(app)
	routes.InitRest(app)

	log.Fatal(app.Listen(":3000"))
}
