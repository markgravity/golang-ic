-- +goose Up
-- +goose StatementBegin
CREATE TABLE "public"."users" (
    id uuid DEFAULT uuid_generate_v1(),
    "email" text,
    "encrypted_password" text,
    "created_at" timestamptz,
    "updated_at" timestamptz,
    PRIMARY KEY ("id")
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE "public"."users";
-- +goose StatementEnd
