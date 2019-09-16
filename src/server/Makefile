PORT := 8888
HOST := localhost

run:
	go run api/main.go

req-books-post:
	curl -v -XPOST $(HOST):$(PORT)/books -d '{"title": "hoge", "author": "hogeman", "price": 500, "img_url":"http://hogehoe.com", "rakuten_page_url":"http://rakutenhogehoge","user_id":1}'

req-community-post:
	curl -v -XPOST $(HOST):$(PORT)/community -d '{"user_ids":[1,2,3,4,5], "community_id":1 }'
