package main

import (
	"analytic/entities"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"time"
)

func main() {


	app := fiber.New()
	// Or extend your config for customization
	app.Use(limiter.New(limiter.Config{
		Max:          5,
		Expiration:     10 * time.Second,
		KeyGenerator:          func(c *fiber.Ctx) string {
			return c.Get("x-forwarded-for")
		},
		LimitReached: func(c *fiber.Ctx) error {
			return c.SendFile("./toofast.html")
		},
	}))
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString(entities.Test())
	})
	// // Default middleware config
	// app.Use(limiter.New())


	app.Listen(":3000")

}
