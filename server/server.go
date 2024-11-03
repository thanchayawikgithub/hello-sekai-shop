package server

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/thanchayawikgithub/hello-sekai-shop/config"
	"github.com/thanchayawikgithub/hello-sekai-shop/modules/middleware/middlewareHandler"
	"github.com/thanchayawikgithub/hello-sekai-shop/modules/middleware/middlewareRepository"
	"github.com/thanchayawikgithub/hello-sekai-shop/modules/middleware/middlewareService"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type (
	server struct {
		app    *echo.Echo
		db     *mongo.Client
		config *config.Config
		mw     middlewareHandler.MiddlewareHandler
	}
)

func newMiddleware(config *config.Config) middlewareHandler.MiddlewareHandler {
	mwRepo := middlewareRepository.NewMiddlewareRepository()
	mwService := middlewareService.NewMiddlewareService(mwRepo)
	return middlewareHandler.NewMiddlewareHandler(mwService, config)
}

func (s *server) gracefulShutdown(ctx context.Context, quit <-chan os.Signal) {
	log.Printf("Start service: %s", s.config.App.Name)
	<-quit
	log.Printf("Shutting down service: %s", s.config.App.Name)

	ctxWithTimeOut, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	if err := s.app.Shutdown(ctxWithTimeOut); err != nil {
		log.Fatalf("Error: %v", err)
	}
}

func (s *server) httpListening() {
	if err := s.app.Start(s.config.App.URL); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Error: %v", err)
	}
}

func Start(ctx context.Context, config *config.Config, db *mongo.Client) {
	s := &server{
		app:    echo.New(),
		db:     db,
		config: config,
		mw:     newMiddleware(config),
	}

	s.app.Use(middleware.TimeoutWithConfig(middleware.TimeoutConfig{
		Skipper:      middleware.DefaultSkipper,
		ErrorMessage: "Error: Request timeout",
		Timeout:      30 * time.Second,
	}))

	s.app.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		Skipper:      middleware.DefaultSkipper,
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.PATCH, echo.DELETE},
	}))

	s.app.Use(middleware.BodyLimit("10MB"))

	s.app.Use(middleware.Logger())

	switch s.config.App.Name {
	case "auth":
		s.authServer()
	case "player":
		s.playerServer()
	case "item":
		s.itemServer()
	case "inventory":
		s.inventoryServer()
	case "payment":
		s.paymentServer()
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	go s.gracefulShutdown(ctx, quit)

	s.httpListening()

}
