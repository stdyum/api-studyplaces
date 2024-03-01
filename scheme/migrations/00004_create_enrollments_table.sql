-- +goose Up
-- +goose StatementBegin

CREATE TABLE IF NOT EXISTS enrollments
(
    id             uuid primary key not null default uuid_generate_v4(),
    user_id        uuid             not null,
    study_place_id uuid             not null,
    user_name      varchar          not null,
    role           role             not null,
    type_id        uuid             not null,
    permissions    varchar[]        null,
    accepted       bool             not null default false,
    blocked        bool             not null default false,


    CONSTRAINT fk_study_place FOREIGN KEY (study_place_id)
        REFERENCES study_places (id),

    UNIQUE (user_id, study_place_id)
);

CALL register_updated_at_created_at_columns('enrollments')

-- +goose StatementEnd
