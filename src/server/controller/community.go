package controller_user

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/k-ueki/app2/src/server/model"
	"github.com/k-ueki/app2/src/server/service"
)

var userId int64 = 6

func (u *DBHandler) CreateCommunity(w http.ResponseWriter, r *http.Request) (int, interface{}, error) {
	req := &model.Community{}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return http.StatusBadRequest, nil, err
	}

	CommunityService := service.NewCommunityService(u.DB)
	com, err := CommunityService.Register(userId, req)
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}
	fmt.Println("com", com)

	return http.StatusOK, nil, nil
}
