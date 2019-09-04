-- +goose Up
CREATE TABLE books (
  id int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  title VARCHAR(255) NOT NULL,
  author VARCHAR(255) NOT NULL,
  price INT(11) UNSIGNED NOT NULL,
  img_url VARCHAR(255),
  page_url VARCHAR(255),
  PRIMARY KEY (id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- +goose Down
DROP TABLE books;
