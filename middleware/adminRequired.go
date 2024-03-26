package middleware

import (
	"net/http"

	"github.com/gofiber/fiber/v3"
)



func AdminMiddleware(c fiber.Ctx)error{
	if c.Locals("isAdmin")!=true{
		return c.Status(http.StatusForbidden).Redirect().To("/")
	}
	return c.Next()
}
