
-- +goose Up
INSERT INTO communities(
	name,
	pass
) VALUES (
	"hoge_community",
	"hogehoge"
);

-- +goose Down
