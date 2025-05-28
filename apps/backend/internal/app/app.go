// Package app configures and runs application.
package app

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/terrnit/rebound/backend/config"
	"github.com/terrnit/rebound/backend/internal/controller/router"
	pgrepo "github.com/terrnit/rebound/backend/internal/repository/postgres"
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
	userRepo := pgrepo.NewUserRepository(pg)
	nutritionRepo := pgrepo.NewNutritionRepository(pg)
	mealRepo := pgrepo.NewMealRepository(pg)
	productRepo := pgrepo.NewProductRepository(pg)
	authRepo := pgrepo.NewAuthRepository(pg)

	// Initialize use cases
	userUC := usecase.NewUserUseCase(userRepo)
	nutritionUC := usecase.NewNutritionUseCase(nutritionRepo)
	mealUC := usecase.NewMealUseCase(mealRepo, nutritionRepo)
	productUC := usecase.NewProductUseCase(productRepo, nutritionRepo)
	authUC := usecase.NewAuthUseCase(authRepo, userRepo)

	// HTTP Server
	httpServer := httpserver.New(httpserver.Port(cfg.HTTP.Port), httpserver.Prefork(cfg.HTTP.UsePreforkMode))

	router.NewRouter(
		httpServer.App,
		userUC,
		nutritionUC,
		mealUC,
		productUC,
		authUC,
		l,
	)

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
