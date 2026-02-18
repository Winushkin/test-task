// Package repository contains db operations
package repository

import (
	"context"
	_ "embed"
	"fmt"

	"file-manager/internal/entities"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

var (
	//go:embed sql/get_processed_files.sql
	getProcessedFileQuery string

	//go:embed sql/insert_processed_file.sql
	insertProcessedFileQuery string

	//go:embed sql/insert_record.sql
	insertRecordQuery string

	//go:embed sql/insert_log.sql
	inserLogQuery string

	//go:embed sql/get_records.sql
	getRecordsQuery string

	//go:embed sql/count_records.sql
	countRecordsQuery string
)

type Postgres struct {
	client *pgxpool.Pool
}

func New(pgClient *pgxpool.Pool) *Postgres {
	return &Postgres{
		client: pgClient,
	}
}

func (repo *Postgres) GetProcessedFiles(ctx context.Context) ([]string, error) {
	rows, err := repo.client.Query(ctx, getProcessedFileQuery)
	if err != nil {
		return nil, fmt.Errorf("failed to get processed files: %w", err)
	}
	defer rows.Close()

	filenames := make([]string, 0)
	for rows.Next() {
		var filename string
		err := rows.Scan(&filename)
		if err != nil {
			return nil, fmt.Errorf("failed to parse row: %w", err)
		}
		filenames = append(filenames, filename)
	}

	return filenames, nil
}

func (repo *Postgres) InsertProcessedFile(ctx context.Context, filename string) (int, error) {
	var fileID int
	err := repo.client.QueryRow(ctx, insertProcessedFileQuery, filename).Scan(&fileID)
	if err != nil {
		return -1, fmt.Errorf("failed to insert processed file: %w", err)
	}
	return fileID, nil
}

func (repo *Postgres) InsertRecord(ctx context.Context, record entities.Record) error {
	args := pgx.NamedArgs{
		"num":           record.Number,
		"mqtt":          record.Mqtt,
		"inv_id":        record.InvID,
		"unit_guid":     record.UnitGUID,
		"message_id":    record.MessageID,
		"message_text":  record.MessageText,
		"context":       record.Context,
		"message_class": record.MessageClass,
		"message_level": record.MessageLevel,
		"area":          record.Area,
		"var_addr":      record.VarAddress,
		"block_sign":    record.Block,
		"message_type":  record.MessageType,
		"bit_number":    record.BitNumber,
		"invert_bit":    record.InvertBit,
		"file_id":       record.FileID,
	}

	_, err := repo.client.Exec(ctx, insertRecordQuery, args)
	if err != nil {
		return fmt.Errorf("failed to insert record: %w", err)
	}
	return nil
}

func (repo *Postgres) InsertLog(ctx context.Context, log entities.Log) error {
	args := pgx.NamedArgs{
		"level":   log.Level,
		"message": log.Message,
	}
	_, err := repo.client.Exec(ctx, inserLogQuery, args)
	if err != nil {
		return fmt.Errorf("failed to insert log: %w", err)
	}
	return nil
}

func (repo *Postgres) GetRecordsWithOffset(ctx context.Context, limit, offset int) ([]entities.Record, error) {
	args := pgx.NamedArgs{
		"limit":  limit,
		"offset": offset,
	}
	rows, err := repo.client.Query(ctx, getRecordsQuery, args)
	if err != nil {
		return nil, fmt.Errorf("failed to get records, limit: %d, offset: %d, err: %w", limit, offset, err)
	}
	return pgx.CollectRows(rows, pgx.RowToStructByName[entities.Record])
}

func (repo *Postgres) CountRecords(ctx context.Context) (int, error) {
	var amount int
	err := repo.client.QueryRow(ctx, countRecordsQuery).Scan(&amount)
	if err != nil {
		return 0, fmt.Errorf("failed to count records: %w", err)
	}
	return amount, nil
}
