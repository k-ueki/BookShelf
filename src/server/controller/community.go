package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/k-ueki/app2/src/server/model"
	"github.com/k-ueki/app2/src/server/service"
)

var userId int64 = 6

type Community struct {
	DB *DBHandler
}

func (u *DBHandler) CreateCommunity(w http.ResponseWriter, r *http.Request) (int, interface{}, error) {
	req := &model.Community{}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return http.StatusBadRequest, nil, err
	}

	CommunityService := service.NewCommunityService(u.DB)
	_, err := CommunityService.Register(userId, req)
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	return http.StatusOK, nil, nil
}

func (u *DBHandler) GetTheCommunityInfo(w http.ResponseWriter, r *http.Request) (int, interface{}, error) {
	vars := mux.Vars(r)
	tmp := vars["com_id"]
	cid, err := strconv.Atoi(tmp)
	if err != nil {
		return http.StatusBadRequest, nil, err
	}

	CommunityService := service.NewCommunityService(u.DB)
	res, err := CommunityService.GetAll(cid)
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	return http.StatusOK, res, nil
}

func (u *DBHandler) AddMembers(w http.ResponseWriter, r *http.Request) (int, interface{}, error) {
	type Request struct {
		Ids   []int64 `json:"ids"`
		ComId int64   `json:"com_id"`
	}
	req := Request{}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return http.StatusBadRequest, nil, err
	}
	req.Ids = req.Ids[1:]

	communityService := service.NewCommunityService(u.DB)
	if err := communityService.Create(req.Ids, req.ComId); err != nil {
		return http.StatusInternalServerError, nil, err
	}

	return http.StatusOK, nil, nil
}
