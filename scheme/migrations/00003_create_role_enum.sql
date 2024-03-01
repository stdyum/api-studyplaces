-- +goose Up
-- +goose StatementBegin

CREATE TYPE role AS ENUM ('student', 'teacher', 'stuff', 'admin');

-- +goose StatementEnd
