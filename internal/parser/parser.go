// Package parser user for parse .tsv files as a part of files' processing
package parser

import (
	"encoding/csv"
	"file-manager/internal/entities"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
)

func readTSVFile(filename string) ([][]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("Open: failed to open file %w", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.Comma = '\t'

	fileData, err := reader.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("ReadAll: failed to read tsv file: %w", err)
	}

	return fileData, nil
}

func ParseTSVFile(filename string) ([]entities.Record, error) {
	ext := filepath.Ext(filename)
	if ext != ".tsv" {
		return nil, fmt.Errorf("Ext: incorrect file extension: expected: tsv, got: %s", ext)
	}

	fileArr, err := readTSVFile(filename)
	if err != nil {
		return nil, fmt.Errorf("readTSVFile: %w", err)
	}

	recordArr := make([]entities.Record, 0)

	for _, line := range fileArr {
		if _, err := strconv.Atoi(line[0]); err != nil {
			continue // скип первых строк с неймами атрибутов
		}
		record := entities.Record{
			Number:       line[0],
			Mqtt:         line[1],
			InvID:        line[2],
			UnitGUID:     line[3],
			MessageID:    line[4],
			MessageText:  line[5],
			Context:      line[6],
			MessageClass: line[7],
			MessageLevel: line[8],
			Area:         line[9],
			VarAddress:   line[10],
			Block:        line[11],
			MessageType:  line[12],
			BitNumber:    line[13],
			InvertBit:    line[14],
		}
		recordArr = append(recordArr, record)
	}
	return recordArr, nil
}
