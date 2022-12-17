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
		arguments     []interface{}
		result        buildQueryResult
	)

	query.WriteString(req.query)

	sortDirection = req.monsterListReq.SortDirection

	arguments = append(arguments, strconv.FormatInt(req.userID, 10))

	if req.monsterListReq.Query != "" {
		arguments = append(arguments, req.monsterListReq.Query)

		where = append(where, fmt.Sprintf("name = %s", "$"+strconv.Itoa(len(arguments))))
	}

	if len(req.monsterListReq.QueryType) > 0 {
		arguments = append(arguments, pq.Array(req.monsterListReq.QueryType))

		where = append(where, fmt.Sprintf("type = ANY(%s)", "$"+strconv.Itoa(len(arguments))))
	}

	query.WriteString(" WHERE " + strings.Join(where, " AND "))

	if _, ok := model.SUPPORTED_SORT_DIRECTION[sortDirection]; ok {
		arguments = append(arguments, model.MAP_SORT_DIRECTION[sortDirection])
		query.WriteString(fmt.Sprintf(" ORDER BY %s %s", req.monsterListReq.Sort, "$"+strconv.Itoa(len(arguments))))
	}

	result.args = arguments
	result.query = query.String()

	return result, nil
}
