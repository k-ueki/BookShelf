
-- +goose Up
CREATE TABLE user_id (
  user_id int(10) UNSIGNED NOT NULL,
  disp_name VARCHAR(255) NOT NULL,
  FOREIGN KEY(user_id) REFERENCES users(id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- +goose Down
DROP TABLE user_id;
