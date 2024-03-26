package middleware

import (
	"log"
	"net/http"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/session"
)

func AuthMiddleware(store *session.Store) fiber.Handler {
	return func(c fiber.Ctx) error {
		sess, err := store.Get(c)
		if err != nil {
			return c.Status(http.StatusUnauthorized).Redirect().To("/login")
		}

		if len(sess.Keys()) > 0 {
			c.Locals("uid", sess.Get("uid"))
			log.Println("mid ",sess.Get("isAdmin"))
			c.Locals("isAdmin", sess.Get("isAdmin"))
			return c.Next()
		} else {
			return c.Status(http.StatusUnauthorized).Redirect().To("/login")
		}
	}
}
