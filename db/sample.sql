SELECT id,name,photo_url FROM users WHERE firebase_uid="hoge"

SELECT 
	books.id,
	books.title,
	books.author,
	books.price,
	books.img_url,
	books.page_url
FROM books
INNER JOIN user_book
ON user_book.book_id=books.id
WHERE user_book.user_id=1

