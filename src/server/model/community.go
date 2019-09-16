package model

type Community struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

type PostCommunityRequest struct {
	UserIds     []int64 `json:"user_ids"`
	CommunityId int64   `json:"community_id"`
}

type CommunityParams struct {
	Uid int64 `db:"user_id"`
	Cid int64 `db:"community_id"`
}
