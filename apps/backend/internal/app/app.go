// Package app configures and runs application.
package app

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/terrnit/rebound/backend/config"
	"github.com/terrnit/rebound/backend/internal/controller/router"
	repo "github.com/terrnit/rebound/backend/internal/repository"
	"github.com/terrnit/rebound/backend/internal/usecase"
	"github.com/terrnit/rebound/backend/pkg/httpserver"
	"github.com/terrnit/rebound/backend/pkg/logger"
	pgpkg "github.com/terrnit/rebound/backend/pkg/postgres"
)

// Run creates objects via constructors.
func Run(cfg *config.Config) {
	l := logger.New(cfg.Log.Level)

	// Repository
	pg, err := pgpkg.New(cfg.PG.URL, pgpkg.MaxPoolSize(cfg.PG.PoolMax))
	if err != nil {
		l.Fatal(fmt.Errorf("app - Run - postgres.New: %w", err))
	}
	defer pg.Close()

	// Initialize repositories
	foodItemRepo := repo.NewFoodItemRepository(pg)
	userRepo := repo.NewUserRepository(pg)

	// Initialize use cases
	foodItemUC := usecase.NewFoodItemUseCase(foodItemRepo, *&usecase.Config{MaxPageSize: 100, DefaultPageSize: 10})
	userUC := usecase.NewUserUseCase(userRepo, *&usecase.UserConfig{MaxPageSize: 100, DefaultPageSize: 10})

	// HTTP Server
	httpServer := httpserver.New(httpserver.Port(cfg.HTTP.Port), httpserver.Prefork(cfg.HTTP.UsePreforkMode))

	router.NewRouter(httpServer.App, userUC, foodItemUC, l)

	// Start servers
	httpServer.Start()

	// Waiting signal
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		l.Info("app - Run - signal: %s", s.String())
	case err = <-httpServer.Notify():
		l.Error(fmt.Errorf("app - Run - httpServer.Notify: %w", err))
	}

	// Shutdown
	err = httpServer.Shutdown()
	if err != nil {
		l.Error(fmt.Errorf("app - Run - httpServer.Shutdown: %w", err))
	}
}
