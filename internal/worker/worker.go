// Package worker used for process tsv files
package worker

import (
	"context"
	"file-manager/internal/config"
	"file-manager/internal/logger"
	"file-manager/internal/parser"
	"file-manager/internal/report"
	"file-manager/internal/repository"
	"fmt"
	"sync"
)

func Work(
	ctx context.Context,
	filePipe <-chan string,
	wg *sync.WaitGroup,
	pgRepo *repository.Postgres,
) {
	defer wg.Done()

	cfg := ctx.Value(logger.ConfigKey).(*config.Config)
	appLogger := ctx.Value(logger.LoggerKey).(*logger.Logger)

	for filename := range filePipe {
		records, err := parser.ParseTSVFile(fmt.Sprintf("%s/%s", cfg.TSVDirPath, filename))
		if err != nil {
			appLogger.Fatal(ctx, fmt.Sprintf("ParseTSVFile: %v", err))
		}

		report.CreateReportsFromFile(records, cfg.ReportsDirPath)

		fileID, err := pgRepo.InsertProcessedFile(ctx, filename)
		if err != nil {
			appLogger.Fatal(ctx, fmt.Sprintf("InsertProcessedFile: %v", err))

		}

		for _, record := range records {
			record.FileID = fileID
			if err = pgRepo.InsertRecord(ctx, record); err != nil {
				appLogger.Fatal(ctx, fmt.Sprintf("InsertRecord: %v", err))
			}
		}
	}
}
