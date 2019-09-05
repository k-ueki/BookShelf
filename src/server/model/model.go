package model

// type Community struct {
// 	ID    int
// 	Name  string
// 	Pass  string
// 	Users string
// 	User1 string
// 	User2 string
// 	User3 string
// 	User4 string
// 	User5 string
// 	User6 string
// 	User7 string
// 	User8 string
// 	User9 string
// }

type User struct {
	Id       int64  `db:"id" json:"id"`
	Uid      string `db:"firebase_uid" json:"uid"`
	Name     string `db:"name" json:"name"`
	PhotoUrl string `db:"photo_url" json:"photo_url"`
}

type UserResp struct {
	Id    int64  `json:"id"`
	Name  string `json:"name"`
	Books []Book `json:"books"`
	//File     string `json:file`
}

type Book struct {
	Id      int64   `db:"id" json:"id"`
	Title   string  `db:"title" json:"title"`
	Author  string  `db:"author" json:"author"`
	Price   int64   `db:"price" json:"price"`
	ImgUrl  string  `db:"img_url" json:"img_url"`
	PageUrl *string `db:"page_url" json:"page_url"`
}

type Item struct {
	Title     string `json:"title"`
	Author    string `json:"author"`
	Isbn      string `json:"isbn"`
	SalesDate string `json:"salesDate"`
	ItemPrice int64  `json:"itemPrice"`
	ItemUrl   string `json:"itemUrl"`
	ImageUrl  string `json:"mediumImageUrl"`
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
	PageUrl string `json:"page_url"`
	UserId  int64  `json:"user_id"`
}

type Community struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

// func (u *User) Insert(db *sql.DB) (*User, error) {
// 	_, err := db.Exec("insert into users (username,user_id,email,password) values (?,?,?,?)", u.Name, u.UserId, u.Email, u.Password)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &User{
// 		Name:     u.Name,
// 		UserId:   u.UserId,
// 		Password: u.Password,
// 		Email:    u.Email,
// 		//File:u.File
// 	}, nil
// }
// func (b *Book) InsertBook(db *sql.DB) (*Book, error) {
// 	_, err := db.Exec("insert into books (title,author,price,img_url,rakuten_page_url,user_id) values (?,?,?,?,?,?)", b.Title, b.Author, b.Price, b.ImgUrl, b.RakutenPageUrl, b.User_id)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return b, nil
// }
// func (b *Book) DeleteBook(db *sql.DB) error {
// 	_, err := db.Exec("delete from books where id=?", b.Id)
// 	return err
// }
// func (u *User) GetInfoByEmailPass(db *sql.DB) (*User, error) {
// 	usr := &User{}
// 	if err := db.QueryRow(`select id,username,email,password from users where email=? AND password=?`, u.Email, u.Password).Scan(&usr.ID, &usr.Name, &usr.Email, &usr.Password); err != nil {
// 		return nil, err
// 	}
// 	return usr, nil
// }
// func (u *User) SelectById(db *sql.DB) (*User, error) {
// 	usr := &User{
// 		ID: u.ID,
// 	}
// 	if err := db.QueryRow(`select username,user_id,email from users where id=?`, u.ID).Scan(&usr.Name, &usr.UserId, &usr.Email); err != nil {
// 		return nil, err
// 	}
// 	return usr, nil
// }
// func (u *User) SelectIdByNameANDPass(db *sql.DB) (*User, error) {
// 	usr := &User{}
// 	if err := db.QueryRow(`select id from users where username=? AND password=?`, u.Name, u.Password).Scan(&usr.ID); err != nil {
// 		return nil, err
// 	}
// 	return usr, nil
// }
//
// func (u *User) SelectUserIDByName(db *sql.DB) (*User, error) {
// 	usr := &User{}
// 	if err := db.QueryRow(`select user_id from users where username=?`, u.Name).Scan(&usr.UserId); err != nil {
// 		return nil, err
// 	}
// 	fmt.Println("OOJO", usr)
// 	return usr, nil
// }
// func (u *User) SelectPersonalBooks(db *sql.DB) ([]Book, error) {
// 	books := []Book{}
// 	bo := Book{}
//
// 	rows, _ := db.Query(`select id,title,author,price,img_url,rakuten_page_url from books where user_id=?`, u.ID)
//
// 	for rows.Next() {
// 		err := rows.Scan(&bo.Id, &bo.Title, &bo.Author, &bo.Price, &bo.ImgUrl, &bo.RakutenPageUrl)
// 		if err != nil {
// 			fmt.Println("Error!", err)
// 		}
// 		books = append(books, bo)
// 	}
// 	fmt.Println("BOOKS", books)
// 	fmt.Println("BO", bo)
// 	return books, nil
// }
// func (u *User) SelectAllPerson(db *sql.DB) ([]User, error) {
// 	Users := []User{}
// 	us := User{}
//
// 	rows, err := db.Query(`select id,username from users`)
// 	if err != nil {
// 		return nil, err
// 	}
//
// 	for rows.Next() {
// 		err := rows.Scan(&us.ID, &us.Name)
// 		if err != nil {
// 			return nil, err
// 		}
// 		Users = append(Users, us)
// 	}
//
// 	return Users, nil
// }
//
// //過去の登録者にuserid,emailがいずれかが重複する人がいないかどうか
// // bool : false=登録不可 , true:登録可
// // error : エラー
// // string : 登録可であれば"" , 登録不可の時は不可の理由("email"であればemailがすでに登録されているもの)
// func (u *User) CheckPreDB(db *sql.DB) (bool, error, string) {
// 	userid := u.UserId
// 	email := u.Email
// 	var count int
// 	var witch string
//
// 	check := func(column string, str string) (int, error) {
// 		rows, err := db.Query(`SELECT COUNT(id) FROM users WHERE `+column+`=?`, str)
// 		if err != nil {
// 			fmt.Println(err)
// 			return 0, err
// 		}
// 		defer rows.Close()
//
// 		for rows.Next() {
// 			_ = rows.Scan(&count)
// 		}
// 		if count >= 1 {
// 			return count, nil
// 		}
// 		return 0, nil
// 	}
//
// 	checkuserid, err := check("user_id", userid)
// 	if err != nil {
// 		return false, err, witch
// 	}
//
// 	checkemail, err := check("email", email)
// 	if err != nil {
// 		return false, err, witch
// 	}
//
// 	fmt.Println(checkuserid, checkemail)
// 	if checkuserid == 0 && checkemail != 0 {
// 		witch = "email"
// 		return false, nil, witch
// 	} else if checkuserid != 0 && checkemail == 0 {
// 		witch = "userid"
// 		return false, nil, witch
// 	} else if checkuserid == 0 && checkemail == 0 {
// 		return true, nil, witch
// 	}
// 	witch = "useridemail"
// 	return false, nil, witch
// }
//
// //func (c *Community) Resister(db *sql.DB) error {
// //	fmt.Println("C", c)
// //
// //	var err error
// //	switch c.Users {
// //	case "1":
// //		_, err = db.Exec("insert into community (com_name,com_pass,users,user1) value (?,?,?,?)", c.Name, c.Pass, c.Users, c.User1)
// //	case "2":
// //		_, err = db.Exec("insert into community (com_name,com_pass,users,user1,user2) value (?,?,?,?)", c.Name, c.Pass, c.Users, c.User1, c.User2)
// //	}
// //	if err != nil {
// //		return err
// //	}
// //	return nil
// //}
