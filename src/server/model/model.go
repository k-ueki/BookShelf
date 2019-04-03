package model

import (
	"database/sql"
	"fmt"
)

type User struct {
	ID       int    `json:id`
	Name     string `json:name`
	Password string `json:pass`
	Email    string `json:email`
	//File     string `json:file`
}

func (u *User) Insert(db *sql.DB) (*User, error) {
	res, err := db.Exec("insert into users (username,email,password) values (?,?,?)", u.Name, u.Email, u.Password)
	fmt.Println("RES", res)
	fmt.Println(err)
	defer db.Close()
	if err != nil {
		return nil, err
	}
	return &User{
		Name:     u.Name,
		Password: u.Password,
		Email:    u.Email,
		//File:u.File
	}, nil
}
func (u *User) Check(db *sql.DB) (*User, error) {
	fmt.Println("JKJK", u)
	fmt.Println(u.Password)

	usr := &User{}
	if err := db.QueryRow(`select id,username,email,password from users where email=? AND password=?`, u.Email, u.Password).Scan(&usr.ID, &usr.Name, &usr.Email, &usr.Password); err != nil {
		return nil, err
	}
	defer db.Close()
	fmt.Println("JJJ", usr)
	return usr, nil
}
