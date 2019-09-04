
-- +goose Up
INSERT INTO user_community(
	user_id,
	community_id
) VALUES (
	1,
	1
);

-- +goose Down
