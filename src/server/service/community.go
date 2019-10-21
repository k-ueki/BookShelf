package service

import (
	"fmt"

	"github.com/gocraft/dbr"
	"github.com/jmoiron/sqlx"
	"github.com/k-ueki/app2/src/server/model"
	"github.com/k-ueki/app2/src/server/repository"
)

type Community struct {
	DB *sqlx.DB
}

func NewCommunityService(db *sqlx.DB) *Community {
	return &Community{db}
}

func (c *Community) Create(uids []int64, cid int64) error {
	coms := make([]model.CommunityParams, 0)

	for _, v := range uids {
		tmp := model.CommunityParams{
			Uid: v,
			Cid: cid,
		}
		coms = append(coms, tmp)
	}

	conn, _ := dbr.Open("mysql", "root:@/uchihon", nil)
	sess := conn.NewSession(nil)

	_, err := repository.CreateCommunity(c.DB, sess, coms)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

func (c *Community) GetCommunitiesByUid(uid string) ([]model.Community, error) {
	communities, err := repository.GetAllByUid(c.DB, uid)
	if err != nil {
		fmt.Println(err)
		// return res, err
	}

	// marUsr, _ := json.Marshal(res)

	// fmt.Fprintf(w, string(marUsr))
	return *communities, nil
}

func (c *Community) Register(userId int64, req *model.Community) (*model.Community, error) {
	result, err := repository.CreateNewCommunity(c.DB, req)
	if err != nil {
		return nil, err
	}

	comid, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	sl := []int64{userId}
	if err := c.Create(sl, comid); err != nil {
		return nil, err
	}

	res := req
	res.Id = comid

	return res, nil
}
