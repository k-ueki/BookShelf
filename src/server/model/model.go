package model

import (
	"database/sql"
	"fmt"
)

type User struct {
	Name     string `json:name`
	Password string `json:pass`
	Email    string `json:email`
	//File     string `json:file`
}

func (u *User) Insert(db *sql.DB) (*User, error) {
	res, err := db.Exec("insert into users (username,email,password) values (?,?,?)", u.Name, u.Email, u.Password)
	if err != nil {
		return nil, err
	}
	fmt.Println(res)
	defer db.Close()
	return &User{
		Name:     u.Name,
		Password: u.Password,
		Email:    u.Email,
		//File:u.File
	}, nil
}
