package pkg

import (
	"github.com/gofiber/fiber/v2"
)

var (
	sCtx *ServerEnvironment
)

func SetContext(ctx *fiber.Ctx) {
	sCtx = &ServerEnvironment{
		Hostname: ctx.Hostname(),
		Url:      ctx.OriginalURL(),
		Method:   ctx.Method(),
	}
}

func GetCtx() *ServerEnvironment {
	defer func() {
		sCtx = nil
	}()
	return sCtx
}
