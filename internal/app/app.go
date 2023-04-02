package app

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	config "github.com/audryus/boleto-palm-tree/configs"
	grpcrouter "github.com/audryus/boleto-palm-tree/internal/controller/grpc"
	"github.com/audryus/boleto-palm-tree/internal/usecase"
	"github.com/audryus/boleto-palm-tree/internal/usecase/repo"
	"github.com/audryus/boleto-palm-tree/pkg/grpcserver"
	"github.com/audryus/boleto-palm-tree/pkg/logger"
	"github.com/audryus/boleto-palm-tree/pkg/postgres"
)

func Run(cfg *config.Config) {
	l := logger.New()

	// Repository
	pg, err := postgres.New(cfg.DB.URL, postgres.MaxPoolSize(cfg.DB.PoolMax))
	if err != nil {
		l.Fatal(fmt.Errorf("app - Run - postgres.New: %w", err))
	}
	defer pg.Close()

	router := grpcrouter.NewCityRouter(usecase.NewCityUseCase(
		repo.NewCityRepo(pg),
	))

	server := grpcserver.New(cfg.GRPC.Port, router)

	// Waiting signal
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		l.Info("app - Run - signal: " + s.String())
	}

	server.GracefulStop()
}
