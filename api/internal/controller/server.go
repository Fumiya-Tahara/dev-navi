package controller

import (
	"context"
	"os"
	"os/signal"
	"time"

	"github.com/labstack/echo/v4"
)

func StartServer(e *echo.Echo, si ServerInterface) {
	RegisterHandlers(e, si)

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	go func() {
		if err := e.Start("8080"); err != nil {
			e.Logger.Fatal("shutting down the server: ", err)
		}
	}()

	<-ctx.Done()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}

}
