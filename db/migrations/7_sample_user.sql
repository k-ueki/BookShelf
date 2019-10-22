
-- +goose Up
INSERT INTO users(
	firebase_uid,
	name,
	photo_url
) VALUES (
	"hoge",
	"hoge太郎",
	"http://hogehoge.com"
),(
	"abc",
	"hoge侍",
	"hogehoge"
),(
	"def",
	"ho",
	"hoho"
),(
	"ghi",
	"piyo",
	"pityopiyo"
),(
	"nil",
	"nil",
	"nilnil"
);

-- +goose Down
