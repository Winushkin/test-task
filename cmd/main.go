package main

import (
	"context"
	_ "file-manager/docs"
	"file-manager/internal/config"
	"file-manager/internal/handlers"
	"file-manager/internal/logger"
	"file-manager/internal/poller"
	"file-manager/internal/postgres"
	"file-manager/internal/repository"
	"file-manager/internal/worker"
	"file-manager/migrator"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"sync"
	"syscall"
	"time"

	httpSwagger "github.com/swaggo/http-swagger"
)

func main() {
	rootCtx := context.Background()
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatal("failed to get config: ", err)
		return
	}

	rootCtx = context.WithValue(rootCtx, logger.ConfigKey, cfg)

	pgClient, err := postgres.NewConn(rootCtx, cfg.PostgresCfg)
	if err != nil {
		log.Fatal("failed to connect to postgres db:", err)
	}

	if err = migrator.Up(pgClient); err != nil {
		log.Fatal("failed to up migrations:", err)
	}

	pgRepo := repository.New(pgClient)

	logFilename := strconv.FormatInt(time.Now().UnixNano(), 10) + ".log"
	appLogger := &logger.Logger{
		Repo:        pgRepo,
		LogFilepath: cfg.LogDIRPath + "/" + logFilename,
	}
	rootCtx = context.WithValue(rootCtx, logger.LoggerKey, appLogger)

	// @title           Records Service API
	// @version         1.0
	// @description     API for viewing TSV files
	// @host            localhost:8080
	// @BasePath        /
	handler := handlers.Handler{Ctx: rootCtx, Repo: pgRepo}
	http.HandleFunc("/records", handler.GetRecordsHandler)
	http.Handle("/swagger/", httpSwagger.WrapHandler)

	go func() {
		appLogger.Info(rootCtx, "http server UP")
		if err := http.ListenAndServe(":8080", nil); err != nil {
			appLogger.Fatal(rootCtx, "http server down: "+err.Error())
		}
	}()

	pollerCtx, cancel := context.WithCancel(rootCtx)
	filesQueue := make(chan string)

	go poller.ScanDirectory(pollerCtx, filesQueue, pgRepo)

	var wg sync.WaitGroup
	wg.Add(1)
	go worker.Work(rootCtx, filesQueue, &wg, pgRepo)

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-sigChan
		appLogger.Info(rootCtx, "Завершение...")
		cancel()
	}()

	wg.Wait()
}
