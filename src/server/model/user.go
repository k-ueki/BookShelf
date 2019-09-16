package model

type User struct {
	Id       int64  `db:"id" json:"id"`
	Uid      string `db:"firebase_uid" json:"uid"`
	Name     string `db:"name" json:"name"`
	PhotoUrl string `db:"photo_url" json:"photo_url"`
}

type UserResp struct {
	Id          int64       `json:"id"`
	Name        string      `json:"name"`
	Books       []Book      `json:"books"`
	Communities []Community `json:"communities"`
	//File     string `json:file`
}
