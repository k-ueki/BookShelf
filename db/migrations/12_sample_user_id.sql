
-- +goose Up
INSERT INTO user_id(
	user_id,
	disp_name
)VALUES(
	1,
	"hoge_taro"
),(
	2,
	"hoge_samurai"
),(
	3,
	"def_desu"
),(
	4,
	"piyo"
),(
	5,
	"nil_nil"
);

-- +goose Down

