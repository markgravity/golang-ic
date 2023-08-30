-- +goose Up
-- +goose StatementBegin
CREATE TABLE "keywords" (
     "id" uuid DEFAULT uuid_generate_v1(),
     "user_id" uuid NOT NULL,
     "text" text NOT NULL,
     "status" text NOT NULL DEFAULT 'pending'::text,
     "links_count" int4 NOT NULL,
     "non_adword_links" json,
     "non_adword_links_count" int4 NOT NULL,
     "adword_links" json,
     "adword_links_count" int4 NOT NULL,
     "html_code" text NOT NULL,
     "created_at" timestamptz NOT NULL,
     "updated_at" timestamptz NOT NULL,
     CONSTRAINT "keywords_user_id_fkey" FOREIGN KEY ("user_id") REFERENCES "users"("id"),
     PRIMARY KEY ("id")
);
CREATE INDEX "keywords_text" ON "keywords" USING BTREE ("text");
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE "keywords";
-- +goose StatementEnd
