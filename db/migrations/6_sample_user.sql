
-- +goose Up
INSERT INTO users(
	firebase_uid,
	name,
	photo_url
) VALUES (
	"hoge",
	"hoge太郎",
	"http://hogehoge.com"
);

-- +goose Down
