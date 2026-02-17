package entities

import "time"

type Log struct {
	ID        uint64    `db:"id"`
	Level     string    `db:"level"`
	Message   string    `db:"message"`
	CreatedAt time.Time `db:"created_at"`
}
