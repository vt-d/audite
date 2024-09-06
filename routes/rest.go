package routes

// unfinished

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vt-d/audite/routes/rest"
)

func InitRest(app *fiber.App) {
	app.Get("/version", rest.Version)
}
