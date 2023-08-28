-- +goose Up
-- +goose StatementBegin
CREATE TABLE "users" (
    id uuid DEFAULT uuid_generate_v1(),
    "email" citext UNIQUE,
    "encrypted_password" text,
    "created_at" timestamptz,
    "updated_at" timestamptz,
    PRIMARY KEY ("id")
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE "users";
-- +goose StatementEnd
