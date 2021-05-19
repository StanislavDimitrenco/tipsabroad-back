package main

import (
	"context"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/stripe/stripe-go/v71"
	"github.com/webdelo/tipsabroad-backend/server"
	"os"
)

func Run() {
	stripe.Key = os.Getenv("STRIPE_KEY")
	terminalOsC := make(chan os.Signal)

	ctx := context.Background()

	_ = server.Run(ctx)
	terminationListening(ctx, terminalOsC)
}

func terminationListening(ctx context.Context, terminalOsC chan os.Signal) {
	<-terminalOsC
	// stop web-server
	if webserver, ok := ctx.Value("webserver").(*fiber.App); ok {
		if err := webserver.Shutdown(); err != nil {
			fmt.Println("Can't shutdown server", err)
		} else {
			fmt.Println("Web-Server was successfully stopped")
		}
	}
	os.Exit(1)

}
