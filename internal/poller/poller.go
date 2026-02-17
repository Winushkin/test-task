// Package poller used for polling tsv directory
package poller

import (
	"context"
	"file-manager/internal/config"
	"file-manager/internal/logger"
	"file-manager/internal/repository"
	"fmt"
	"os"
	"path/filepath"
	"slices"
	"time"
)

func filterFiles(entries []os.DirEntry, processedFiles []string) []string {
	files := make([]string, 0)
	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}
		name := entry.Name()
		if filepath.Ext(name) != ".tsv" {
			continue
		}
		if slices.Contains(processedFiles, name) {
			continue
		}
		files = append(files, name)
	}
	return files
}

func ScanDirectory(
	ctx context.Context,
	filesPipe chan<- string,
	pgRepo *repository.Postgres,
) {
	cfg := ctx.Value(logger.ConfigKey).(*config.Config)
	appLogger := ctx.Value(logger.LoggerKey).(*logger.Logger)

	ticker := time.NewTicker(time.Duration(cfg.PollingInterval) * time.Second)
	defer ticker.Stop()
	defer close(filesPipe)

	for {
		select {
		case <-ctx.Done():
			return

		case <-ticker.C:
			entries, err := os.ReadDir(cfg.TSVDirPath)
			if err != nil {
			}

			processedFiles, err := pgRepo.GetProcessedFiles(ctx)
			if err != nil {
				appLogger.Fatal(ctx, fmt.Sprintf("GetProcessedFiles: %v", err))
			}

			newTSVFiles := filterFiles(entries, processedFiles)
			if amount := len(newTSVFiles); amount > 0 {
				appLogger.Info(ctx, fmt.Sprintf("Найдено %d новых файлов", amount))
			}
			for _, filename := range newTSVFiles {
				filesPipe <- filename
			}
		}
	}
}
