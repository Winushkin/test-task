-- +goose Up
-- +goose StatementBegin
CREATE TABLE logs (
    id BIGSERIAL PRIMARY KEY,
    level VARCHAR(10),
    message TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_logs_level on logs(level);
CREATE INDEX idx_logs_created_at on logs(created_at DESC);
-- +goose StatementEnd


-- +goose Down
-- +goose StatementBegin
DROP TABLE logs;
-- +goose StatementEnd
