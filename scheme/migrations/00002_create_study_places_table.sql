-- +goose Up
-- +goose StatementBegin

CREATE TABLE IF NOT EXISTS study_places
(
    id    uuid primary key not null default uuid_generate_v4(),
    title varchar unique   not null
);


CALL register_updated_at_created_at_columns('study_places');

CREATE INDEX IF NOT EXISTS study_places_created_at_idx ON study_places USING hash (created_at);

-- +goose StatementEnd
