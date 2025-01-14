-- +goose Up
-- +goose StatementBegin
ALTER TABLE rooms
    ADD COLUMN description TEXT;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE rooms
    DROP COLUMN description;
-- +goose StatementEnd
