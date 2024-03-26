package routes

import (
	components "github.com/Mohammed785/goBlog/components/post"
	"github.com/Mohammed785/goBlog/controller"
	"github.com/Mohammed785/goBlog/helpers"
	"github.com/Mohammed785/goBlog/middleware"
	"github.com/gofiber/fiber/v3"
)

func SetupPostRoutes(route fiber.Router, postController *controller.PostController, authMiddleware fiber.Handler) {
	route.Use(authMiddleware)
	route.Get("/", postController.List)
	route.Get("/:id<int>", postController.FindOne)
	route.Get("/create", func(c fiber.Ctx) error {
		return helpers.Render(c, components.PostFormPage(nil))
	})
	route.Use(middleware.AdminMiddleware)
	route.Post("/create", postController.Create)
	route.Patch("/:id", postController.Update)
	route.Delete("/:id", postController.Delete)
}
