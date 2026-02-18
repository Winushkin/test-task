// Package logger contains funcs to log app
package logger

import (
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	"file-manager/internal/entities"
	"file-manager/internal/repository"
)

type ctxKey int

const (
	LoggerKey ctxKey = iota
	ConfigKey
)

type Logger struct {
	Repo        *repository.Postgres
	LogFilepath string
}

func (l *Logger) Info(
	ctx context.Context,
	msg string,
) {
	logEntity := entities.Log{
		Level:     "INFO",
		CreatedAt: time.Now(),
		Message:   msg,
	}
	logString := logToString(logEntity)
	if err := l.saveLog(ctx, logString, logEntity); err != nil {
		log.Fatal("failed to save log:", err)
	}
	log.Println(logString)
}

func (l *Logger) Fatal(
	ctx context.Context,
	msg string,
) {
	logEntity := entities.Log{
		Level:     "FATAL",
		CreatedAt: time.Now(),
		Message:   msg,
	}
	logString := logToString(logEntity)
	if err := l.saveLog(ctx, logString, logEntity); err != nil {
		log.Fatal("failed to save log:", err)
	}
	log.Fatal(logString)
}

func (l *Logger) Debug(
	ctx context.Context,
	msg string,
) {
	logEntity := entities.Log{
		Level:     "DEBUG",
		CreatedAt: time.Now(),
		Message:   msg,
	}
	logString := logToString(logEntity)
	if err := l.saveLog(ctx, logString, logEntity); err != nil {
		log.Fatal("failed to save log:", err)
	}
	log.Println(logString)
}

func (l *Logger) saveLog(
	ctx context.Context,
	logStr string,
	logEntity entities.Log,
) error {

	ext := filepath.Ext(l.LogFilepath)
	if ext != ".log" && ext != ".txt" {
		return fmt.Errorf("file extention: expected .log or .txt, got: %s, ", ext)
	}

	file, err := os.OpenFile(l.LogFilepath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	err = l.Repo.InsertLog(ctx, logEntity)
	if err != nil {
		return fmt.Errorf("InsertLog: %w", err)
	}

	err = saveLogtoFile(logStr, file)
	if err != nil {
		return fmt.Errorf("saveLogtoFile: %w", err)
	}

	return nil
}

func saveLogtoFile(logString string, file *os.File) error {
	_, err := file.WriteString(logString + "\n")
	if err != nil {
		return fmt.Errorf("failed to save log to file: %w", err)
	}
	return nil
}

func logToString(log entities.Log) string {
	logString := fmt.Sprintf(
		"%d. %s: %v. Message: %s.",
		log.ID,
		log.Level,
		log.CreatedAt,
		log.Message,
	)
	return logString
}
