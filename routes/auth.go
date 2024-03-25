package routes

import (
	"github.com/Mohammed785/goBlog/components/auth"
	"github.com/Mohammed785/goBlog/controller"
	"github.com/Mohammed785/goBlog/helpers"
	"github.com/gofiber/fiber/v3"
)



func SetupAuthRoutes(route fiber.Router,authController *controller.AuthController){
	route.Get("/login",func(c fiber.Ctx) error {
		return helpers.Render(c, components.AuthPage("login"))
	})
	route.Get("/register",func(c fiber.Ctx) error {
		return helpers.Render(c, components.AuthPage("register"))
	})
	route.Post("/login", authController.Login)
	route.Post("/register", authController.Register)
	route.Post("/logout", authController.Logout)
}
