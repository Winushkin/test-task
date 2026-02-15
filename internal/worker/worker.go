// Package worker used for process tsv files
package worker

import (
	"file-manager/internal/parser"
	"file-manager/internal/report"
	"fmt"
	"log"
	"sync"
)

func mockSaveRecord(_ parser.Record) error {
	// mock funcs for DB record saving
	return nil
}

func mockSaveFile(_ string) error {
	// mock funcs for DB Save Funcs for processed tsv files
	return nil
}

func Work(filePipe <-chan string, reportDir, tsvDir string, wg *sync.WaitGroup) {
	defer wg.Done()
	for filename := range filePipe {
		records, err := parser.ParseTSVFile(fmt.Sprintf("%s/%s", tsvDir, filename))
		if err != nil {
			log.Fatal("ParseTSVFile:", err)
		}

		report.CreateReportsFromFile(records, reportDir)

		for _, record := range records {
			if err = mockSaveRecord(record); err != nil {
				log.Fatal("mockSaveRecord:", err)
			}
		}

		if err = mockSaveFile(filename); err != nil {
			log.Fatal("mockSaveFile:", err)
		}
	}
}
