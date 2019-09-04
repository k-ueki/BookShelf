
-- +goose Up
CREATE TABLE user_book (
  user_id INT(10) UNSIGNED NOT NULL,
  book_id INT(10) UNSIGNED NOT NULL,
  FOREIGN KEY(user_id) REFERENCES users(id),
  FOREIGN KEY (book_id) REFERENCES books(id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- +goose Down
DROP TABLE user_book;
