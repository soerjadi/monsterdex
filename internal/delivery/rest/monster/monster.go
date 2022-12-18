package monster

import (
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
	"github.com/soerjadi/monsterdex/internal/log"
	"github.com/soerjadi/monsterdex/internal/model"
	"github.com/soerjadi/monsterdex/internal/model/constant"
)

func (h *Handler) GetList(w http.ResponseWriter, r *http.Request) (interface{}, error) {
	param := r.URL.Query()
	var (
		monsterType   []int
		userID        = int64(0)
		sortDirection int
		err           error
		query         *string
	)

	queryParam := param.Get("query")
	queryTypeStr := param.Get("monster_type") // should be []int{}
	sort := param.Get("sort")
	sortDirectionStr := param.Get("sort_direction")

	if queryParam != "" {
		query = &queryParam
	}

	if queryTypeStr != "" {
		queryTypeSlc := strings.Split(queryTypeStr, ",")

		for _, queryType := range queryTypeSlc {
			id, err := strconv.Atoi(queryType)
			if err != nil {
				return nil, errors.New("monster_type should be integer")
			}

			monsterType = append(monsterType, id)
		}
	}

	if sortDirectionStr != "" {
		sortDirection, err = strconv.Atoi(sortDirectionStr)
		if err != nil {
			return nil, errors.New("sort_direction should be integer")
		}
	}

	userIDStr := r.Context().Value(constant.USER_ID_KEY)

	if userIDStr != nil {
		userID = userIDStr.(int64)
	}

	request := model.MonsterListRequest{
		Query:         query,
		QueryType:     &monsterType,
		Sort:          &sort,
		SortDirection: &sortDirection,
	}

	res, err := h.usecase.GetListMonster(r.Context(), request, int64(userID))

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (h *Handler) GetDetailMonster(w http.ResponseWriter, r *http.Request) (interface{}, error) {
	userID := int64(0)

	varParam := mux.Vars(r)
	monsterIDParam := varParam["monster_id"]
	monsterID, err := strconv.ParseInt(monsterIDParam, 10, 32)
	if err != nil {
		return nil, err
	}

	userIDStr := r.Context().Value(constant.USER_ID_KEY)

	if userIDStr != nil {
		userID = userIDStr.(int64)
	}

	res, err := h.usecase.GetDetailMonster(r.Context(), monsterID, userID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("not found")
		}

		return nil, err
	}

	return res, nil
}

func (h *Handler) CapturedMonster(w http.ResponseWriter, r *http.Request) (interface{}, error) {

	dec := json.NewDecoder(r.Body)

	req := model.CaptureMonsterReq{}

	if err := dec.Decode(&req); err != nil {
		log.ErrorWithFields("handler.Monster.CaptureMonster fail to decode request body", log.KV{
			"err": err,
		})
		return nil, err
	}

	varParam := mux.Vars(r)
	monsterIDParam := varParam["monster_id"]
	monsterID, err := strconv.ParseInt(monsterIDParam, 10, 32)
	if err != nil {
		return nil, err
	}

	userID := r.Context().Value(constant.USER_ID_KEY).(int64)

	req.MonsterID = monsterID
	req.UserID = userID

	_, err = h.usecase.CaptureMonster(r.Context(), req)
	if err != nil {
		return nil, err
	}

	return "", nil
}

func (h *Handler) InsertMonster(w http.ResponseWriter, r *http.Request) (interface{}, error) {
	dec := json.NewDecoder(r.Body)

	req := model.MonsterRequest{}

	if err := dec.Decode(&req); err != nil {
		log.ErrorWithFields("handler.Monster.InsertMonster fail to decode request body", log.KV{
			"err": err,
		})

		return nil, err
	}

	res, err := h.usecase.InsertMonster(r.Context(), req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (h *Handler) UpdateMonster(w http.ResponseWriter, r *http.Request) (interface{}, error) {
	dec := json.NewDecoder(r.Body)

	req := model.MonsterRequest{}

	if err := dec.Decode(&req); err != nil {
		log.ErrorWithFields("handler.Monster.InsertMonster fail to decode request body", log.KV{
			"err": err,
		})

		return nil, err
	}

	varParam := mux.Vars(r)
	monsterIDParam := varParam["monster_id"]
	monsterID, err := strconv.ParseInt(monsterIDParam, 10, 32)
	if err != nil {
		return nil, err
	}

	req.ID = monsterID

	err = h.usecase.UpdateMonster(r.Context(), req)
	if err != nil {
		return nil, err
	}

	return "", nil
}

func (h *Handler) DeleteMonster(w http.ResponseWriter, r *http.Request) (interface{}, error) {

	varParam := mux.Vars(r)
	monsterIDParam := varParam["monster_id"]
	monsterID, err := strconv.ParseInt(monsterIDParam, 10, 32)
	if err != nil {
		return nil, err
	}

	err = h.usecase.DeleteMonster(r.Context(), monsterID)
	if err != nil {
		return nil, err
	}

	return "", nil
}
