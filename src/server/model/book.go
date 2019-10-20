package model

type Book struct {
	Id      int64   `db:"id" json:"id"`
	Title   string  `db:"title" json:"title"`
	Author  string  `db:"author" json:"author"`
	Price   int64   `db:"price" json:"price"`
	ImgUrl  string  `db:"img_url" json:"img_url"`
	PageUrl *string `db:"page_url" json:"page_url"`
	Isbn    string  `db:"isbn" json:"isbn"`
}

type Item struct {
	Title     string `json:"title"`
	Author    string `json:"author"`
	Isbn      string `json:"isbn"`
	SalesDate string `json:"salesDate"`
	ItemPrice int64  `json:"itemPrice"`
	ItemUrl   string `json:"itemUrl"`
	ImageUrl  string `json:"largeImageUrl"`
}

type Items struct {
	Item Item `json:"Item"`
}

type Resp struct {
	Items []Items `json:"Items"`
}

type Books struct {
	Books []Items `json:books`
}

type PostBookRequest struct {
	Title   string `json:"title"`
	Author  string `json:"author"`
	Price   int64  `json:"price"`
	ImgUrl  string `json:"img_url"`
	PageUrl string `json:"rakuten_page_url"`
	UserId  int64  `json:"user_id"`
	Isbn    string `json:"isbn"`
}
