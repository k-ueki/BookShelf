package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/k-ueki/app2/src/server/model"
)

func GetAllByUid(db *sqlx.DB, uid string) (*[]model.Community, error) {
	communities := make([]model.Community, 0)
	if err := db.Select(&communities, `
SELECT 
	communities.id,
	communities.name
FROM communities
INNER JOIN user_community
ON user_community.community_id=communities.id
INNER JOIN users
ON user_community.user_id=users.id
WHERE users.firebase_uid=?
	`, uid); err != nil {
		return nil, err
	}
	return &communities, nil
}
