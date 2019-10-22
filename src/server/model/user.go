package model

type User struct {
	Id       int64  `db:"id" json:"id"`
	Uid      string `db:"firebase_uid" json:"uid"`
	Name     string `db:"name" json:"name"`
	PhotoUrl string `db:"photo_url" json:"photo_url"`
	DispName string `db:"disp_name" json:"disp_name"`
}

type UserResp struct {
	Id          int64       `json:"id"`
	Name        string      `json:"name"`
	DispName    string      `json:"disp_name"`
	Books       []Book      `json:"books"`
	Communities []Community `json:"communities"`
	//File     string `json:file`
}

type UserReq struct {
	Uid  string `db:"firebase_uid" json:"uid"`
	Name string `db:"name" json:"name"`
}

type UserDispName struct {
	Uid      int    `db:"user_id"`
	DispName string `db:"disp_name"`
}
