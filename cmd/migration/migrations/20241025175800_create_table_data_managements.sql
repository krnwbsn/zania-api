-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS data_managements (
	"id" SERIAL NOT NULL PRIMARY KEY,
	"type" VARCHAR(255) NOT NULL,
	"title" VARCHAR(255) NOT NULL,
	"position" INT NOT NULL,
	"created_at" TIMESTAMPTZ(6),
	"updated_at" TIMESTAMPTZ(6)
);
-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS data_managements;
-- +goose StatementEnd