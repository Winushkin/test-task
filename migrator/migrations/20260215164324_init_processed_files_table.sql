-- +goose Up
-- +goose StatementBegin
CREATE TABLE processed_files (
    id SERIAL PRIMARY KEY,
    filename TEXT NOT NULL,
    processed_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE processed_files;
-- +goose StatementEnd
