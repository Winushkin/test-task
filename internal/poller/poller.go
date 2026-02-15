// Package poller used for polling tsv directory
package poller

import (
	"context"
	"log"
	"os"
	"path/filepath"
	"slices"
	"time"
)

func mockGetFiles() []string {
	return []string{}
}

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
	dirPath string,
	ctx context.Context,
	filesPipe chan<- string,
	interval int,
) {
	ticker := time.NewTicker(time.Duration(interval) * time.Second)
	defer ticker.Stop()
	defer close(filesPipe)

	for {
		select {
		case <-ctx.Done():
			return

		case <-ticker.C:
			entries, err := os.ReadDir(dirPath)
			if err != nil {
				log.Fatal("failed to scan directory:", err)
			}

			processedFiles := mockGetFiles()
			newTSVFiles := filterFiles(entries, processedFiles)
			if amount := len(newTSVFiles); amount > 0 {
				log.Printf("Найдено %d новых файлов", amount)
			}
			for _, filename := range newTSVFiles {
				filesPipe <- filename
			}
		}
	}
}
