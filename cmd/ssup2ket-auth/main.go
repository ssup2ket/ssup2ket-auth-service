package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"github.com/ssup2ket/ssup2ket-auth-service/internal/config"
	"github.com/ssup2ket/ssup2ket-auth-service/internal/domain"
	"github.com/ssup2ket/ssup2ket-auth-service/internal/server/grpc_server"
	"github.com/ssup2ket/ssup2ket-auth-service/internal/server/http_server"
)

func main() {
	// Get config
	cfg := config.GetConfigs()

	// Init logger
	zerolog.TimestampFieldName = "timestamp"
	log.Logger = log.Logger.With().Caller().Logger()
	if cfg.DeployEnv == config.DeployEnvLocal {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	} else if cfg.DeployEnv == config.DeployEnvDev {
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	} else if cfg.DeployEnv == config.DeployEnvProd {
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	} else {
		log.Fatal().Msg("Wrong deploy env")
	}

	// Print config and starting
	log.Info().Str("config", fmt.Sprintf("%+v", *cfg)).Send()
	log.Info().Msg("Starting ssup2ket auth service...")

	// Init domain
	d, err := domain.New(cfg)
	if err != nil {
		log.Fatal().Msg("Failed to create domain instance")
	}

	// Init and run HTTP server
	httpServer, err := http_server.New(cfg.ServerURL, d)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to create HTTP server")
	}
	log.Info().Msg("Starting HTTP server...")
	httpServer.ListenAndServe()

	// Init and run GRPC server
	grpcServer, err := grpc_server.New(d)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to create GRPC server")
	}
	log.Info().Msg("Starting GRPC server...")
	grpcServer.ListenAndServe()

	// Block until receive a terminal signal
	log.Info().Msg("Waiting a terminal signal to shutdown gracefully")
	termSignal := make(chan os.Signal, 1)
	signal.Notify(termSignal, syscall.SIGTERM, syscall.SIGINT)

	// Receive a terminal signal and shutdown gracefully
	<-termSignal
	log.Info().Msg("Receive a terminal signal and shutdown gracefully")

	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		httpServer.Shutdown()
	}()
	go func() {
		defer wg.Done()
		grpcServer.Shutdown()
	}()
	wg.Wait()
}
