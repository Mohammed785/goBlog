package cmd

import (
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/logger"
	"github.com/gofiber/fiber/v3/middleware/recover"
)

func CreateApp() *fiber.App {
	
	app := fiber.New(fiber.Config{
		CaseSensitive: true,
	})
	app.Use(recover.New())
	app.Use(logger.New())
	
	return app
}
