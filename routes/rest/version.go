package rest

import "github.com/gofiber/fiber/v2"

func Version(ctx *fiber.Ctx) error {
	return ctx.SendString("1.0.0")
}
