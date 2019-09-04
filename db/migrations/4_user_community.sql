
-- +goose Up
CREATE TABLE user_community (
  user_id INT(10) UNSIGNED NOT NULL,
  community_id INT(10) UNSIGNED NOT NULL,
  FOREIGN KEY(user_id) REFERENCES users(id),
  FOREIGN KEY (community_id) REFERENCES communities(id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- +goose Down
DROP TABLE user_community;
