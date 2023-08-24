-- +goose Up
-- +goose StatementBegin
CREATE TABLE "keywords" (
     "id" uuid DEFAULT uuid_generate_v1(),
     "user_id" uuid NOT NULL,
     "keyword" text NOT NULL,
     "status" text NOT NULL DEFAULT 'pending'::text,
     "links_count" int8 NOT NULL,
     "non_adword_links" json,
     "non_adword_links_count" int8 NOT NULL,
     "adword_links" json,
     "adword_links_count" int8 NOT NULL,
     "html_code" text NOT NULL,
     "created_at" timestamptz NOT NULL,
     "updated_at" timestamptz NOT NULL,
     PRIMARY KEY ("id")
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE "keywords";
-- +goose StatementEnd
