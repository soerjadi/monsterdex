package monster

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/lib/pq"
	"github.com/soerjadi/monsterdex/internal/model"
)

func buildMonsterListQuery(ctx context.Context, req buildQueryParam) (buildQueryResult, error) {
	var (
		query         strings.Builder
		where         []string
		sortDirection int
		sortByField   string
		arguments     []interface{}
		result        buildQueryResult
	)

	query.WriteString(req.query)

	sortDirection = model.SORT_DIRECTION_ASC
	sortByField = "id"

	if req.monsterListReq.SortDirection != nil {
		sortDirection = *req.monsterListReq.SortDirection
	}

	if req.monsterListReq.Sort != nil {
		sortByField = *req.monsterListReq.Sort
	}

	arguments = append(arguments, strconv.FormatInt(req.userID, 10))

	if req.monsterListReq.Query != nil {
		arguments = append(arguments, *req.monsterListReq.Query)

		where = append(where, fmt.Sprintf("name = %s", "$"+strconv.Itoa(len(arguments))))
	}

	if req.monsterListReq.QueryType != nil && len(*req.monsterListReq.QueryType) > 0 {
		arguments = append(arguments, pq.Array(*req.monsterListReq.QueryType))

		where = append(where, fmt.Sprintf("type && %s", "$"+strconv.Itoa(len(arguments))))
	}

	if len(where) > 0 {
		query.WriteString(" WHERE " + strings.Join(where, " AND "))
	}

	if _, ok := model.SUPPORTED_SORT_DIRECTION[sortDirection]; ok {
		arguments = append(arguments, model.MAP_SORT_DIRECTION[sortDirection])
		query.WriteString(fmt.Sprintf(" ORDER BY %s %s", sortByField, "$"+strconv.Itoa(len(arguments))))
	}

	result.args = arguments
	result.query = query.String()

	return result, nil
}
