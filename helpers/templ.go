package helpers

import (
	"encoding/json"
	"fmt"

	"github.com/a-h/templ"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/adaptor"
)

func Render(c fiber.Ctx, component templ.Component, options ...func(*templ.ComponentHandler)) error {
	componentHandler := templ.Handler(component)
	for _, o := range options {
		o(componentHandler)
	}
	return adaptor.HTTPHandler(componentHandler)(c)
}

const (
	SUCCESS_MSG = "success"
	ERROR_MSG   = "error"
	WARNING_MSG = "warning"
	INFO_MSG    = "info"
)

func SendMsg(c fiber.Ctx, msgType, msg string) {
	key := fmt.Sprintf("%sMsg", msgType)
	htmxEvent, _ := json.Marshal(map[string]string{key: msg})
	c.Append("HX-Trigger", string(htmxEvent))
}
