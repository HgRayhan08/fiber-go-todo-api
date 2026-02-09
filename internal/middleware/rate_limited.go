package middleware

import (
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/limiter"
)

func RateLimited() fiber.Handler {

	return limiter.New(limiter.Config{
		Max:        5,               // max 100 request
		Expiration: 1 * time.Minute, // per 1 menit
		KeyGenerator: func(c fiber.Ctx) string {
			return c.IP() // berdasarkan IP
		},
		LimitReached: func(c fiber.Ctx) error {
			code := fiber.StatusTooManyRequests
			return c.Status(code).JSON(fiber.Map{
				"code":    code,
				"message": "Terlalu banyak request, silakan coba lagi nanti",
			})
		},
	})

}
