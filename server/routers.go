package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/webdelo/tipsabroad-backend/controllers"
)

func routers(app *fiber.App) {

	app.Post("/create-checkout-session", controllers.CreateCheckoutSession)

}
