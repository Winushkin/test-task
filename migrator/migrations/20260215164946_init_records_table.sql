-- +goose Up
-- +goose StatementBegin
CREATE TABLE records (
    id SERIAL PRIMARY KEY,
    num TEXT,
    mqtt TEXT, 
    inv_id CHAR(8) NOT NULL,
    unit_guid TEXT NOT NULL,
    message_id TEXT,
    message_text TEXT,
    context TEXT,
    message_class VARCHAR(7),
    message_level TEXT,
    area TEXT,
    var_addr TEXT,
    block_sign TEXT,
    message_type TEXT,
    bit_number TEXT,
    invert_bit TEXT,
    file_id INTEGER REFERENCES processed_files(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE records;
-- +goose StatementEnd
