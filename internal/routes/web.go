package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mortazakiani/ticket/internal/api"
)

func SetupRoutes(server *api.Server)  error {
	server.App.Get("/",func (c *fiber.Ctx) error {
		return c.SendString("I'm post request")
	})
	return nil
}