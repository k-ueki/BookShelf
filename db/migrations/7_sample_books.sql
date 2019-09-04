
-- +goose Up
INSERT INTO books(
	title,
	author,
	price,
	img_url
) VALUES (
	"hogeho物語",
	"hoge座衛門",
	3000,
	"http://ahahaha.com"
);

-- +goose Down
