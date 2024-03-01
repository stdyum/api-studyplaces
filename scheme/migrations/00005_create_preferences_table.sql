-- +goose Up
-- +goose StatementBegin

CREATE TABLE IF NOT EXISTS preferences
(
    enrollment_id uuid primary key not null,
    website       json             null default null,
    schedule      json             null default null,
    journal       json             null default null,

    CONSTRAINT fk_enrollment_id FOREIGN KEY (enrollment_id)
        REFERENCES enrollments (id)
);

CALL register_updated_at_created_at_columns('preferences')

-- +goose StatementEnd
