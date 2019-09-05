package service

import (
	"fmt"

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
