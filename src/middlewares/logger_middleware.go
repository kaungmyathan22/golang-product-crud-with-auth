package middlewares

import (
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
)

func LoggerMiddleware(ctx *fiber.Ctx) error {
	start := time.Now()
	clientIP := ctx.IP()
	log.Printf("[%s] %s %s %v", clientIP, ctx.Method(), ctx.Request().URI().FullURI(), time.Since(start))
	return ctx.Next()
}
