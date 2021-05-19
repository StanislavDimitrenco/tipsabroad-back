package server

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func Run(ctx context.Context) context.Context {
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin,X-Requested-With, Content-Type, Accept",
	}))

	app.Use(func(c *fiber.Ctx) error {
		c.Locals("ctx", ctx)
		return c.Next()
	})

	routers(app)

	// start web-server
	go func(app *fiber.App) {
		err := app.Listen(":3000")
		if err != nil {
			panic(err)
		}
	}(app)

	return context.WithValue(ctx, "webserver", app)

}
