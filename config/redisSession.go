package config

import (
	"time"

	"github.com/gofiber/fiber/v3/middleware/session"
	"github.com/gofiber/storage/redis/v3"
)

func NewRedisSessionStorage() *session.Store{
	store := redis.New()
	session := session.New(session.Config{
		Storage:        store,
		Expiration:     12 * time.Hour,
		CookieHTTPOnly: true,
		CookieSecure:   Config("environment") == "production",
		KeyLookup:      "cookie:goBlog_session",
	})
	return session
}
