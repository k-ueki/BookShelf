
-- +goose Up
INSERT INTO user_book(
	user_id,
	book_id
) VALUES (
	1,
	1
);

-- +goose Down
