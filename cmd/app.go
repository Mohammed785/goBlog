package cmd

import (
	"github.com/Mohammed785/goBlog/config"
	"github.com/Mohammed785/goBlog/controller"
	"github.com/Mohammed785/goBlog/database"
	"github.com/Mohammed785/goBlog/database/sqlc"
	"github.com/Mohammed785/goBlog/helpers"
	"github.com/Mohammed785/goBlog/routes"

	"github.com/Mohammed785/goBlog/middleware"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/logger"
	"github.com/gofiber/fiber/v3/middleware/recover"
)

func CreateApp() *fiber.App {
	dbPool := database.ConnectToPostgres()
	queries := sqlc.New(dbPool)
	app := fiber.New(fiber.Config{
		CaseSensitive:     true,
		PassLocalsToViews: true,
	})

	app.Use(recover.New())
	app.Use(logger.New())

	store := config.NewRedisSessionStorage()
	validator := helpers.NewValidator()
	authController := controller.NewAuthController(queries, store, validator)
	postController := controller.NewPostController(queries, validator)
	app.Static("/static", "./assets")
	routes.SetupAuthRoutes(app.Group("/"), authController)
	routes.SetupPostRoutes(app.Group("/post"),postController,middleware.AuthMiddleware(store))
	app.Post("/logout", authController.Logout)
	return app
}
