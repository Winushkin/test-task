package main

import (
	"context"
	"file-manager/internal/config"
	"file-manager/internal/poller"
	"file-manager/internal/worker"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

func main() {
	rootCtx := context.Background()

	pollerCtx, cancel := context.WithCancel(rootCtx)
	filesQueue := make(chan string)

	log.Println("Начало программы")

	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatal("failed to get config: ", err)
		return
	}

	go poller.ScanDirectory(cfg.TSVDirPath, pollerCtx, filesQueue, cfg.PollinInterval)

	var wg sync.WaitGroup
	wg.Add(1)
	go worker.Work(filesQueue, cfg.ReportsDirPath, cfg.TSVDirPath, &wg)

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-sigChan
		log.Println("Завершение...")
		cancel()
	}()

	wg.Wait()
	log.Println("Конец")
}
