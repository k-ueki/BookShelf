UID := demo
PORT := 8888
HOST := localhost
# TOKEN_FILE := .idToken

ARTICLE_ID:=1
ARTICLE_TITLE:=title
ARTICLE_BODY:=body

# req-articles:
# 	curl -v $(HOST):$(PORT)/articles
#
# req-img-pei:
# 	curl -v $(HOST):$(PORT)/img/pei.png
#
# req-articles-get:
# 	curl -v $(HOST):$(PORT)/articles/$(ARTICLE_ID)

req-books-post:
	curl -v -XPOST  $(HOST):$(PORT)/books -d '{"title": "hoge", "author": "hogeman", "price": 500, "img_url":"http://hogehoe.com", "rakuten_page_url":"http://rakutenhogehoge","user_id":1}'

# req-articles-update:
# 	curl -v -XPUT -H "Authorization: Bearer $(shell cat ./$(TOKEN_FILE))" $(HOST):$(PORT)/articles/$(ARTICLE_ID) -d '{"title": "$(ARTICLE_TITLE)", "body": "$(ARTICLE_BODY)"}'
#
# req-articles-delete:
# 	curl -v -XDELETE -H "Authorization: Bearer $(shell cat ./$(TOKEN_FILE))" $(HOST):$(PORT)/articles/$(ARTICLE_ID)
#
# req-public:
# 	curl -v $(HOST):$(PORT)/public
#
# req-private:
# 	curl -v -H "Authorization: Bearer $(shell cat ./$(TOKEN_FILE))" $(HOST):$(PORT)/private
#
# database-init:
# 	make -C ../database init
#
# req-adminfood-post:
# 	curl -v -XPOST -H "Authorization: Bearer $(shell cat ./$(TOKEN_FILE))" $(HOST):$(PORT)/admin/foods -d '{"name": "hoge", "price":100,"amount":5,"until":"2020-09-02 12:00:00","image":["https://hogehoge.com","https://hogeeeeeeeee.com"]}'
#
