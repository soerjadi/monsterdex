package model

const (
	SORT_DIRECTION_ASC  = 1
	SORT_DIRECTION_DESC = 2

	QUERY_TYPE_NAME         = 1
	QUERY_TYPE_MONSTER_TYPE = 2
)

var SUPPORTED_SORT_DIRECTION = map[int]bool{
	SORT_DIRECTION_ASC:  true,
	SORT_DIRECTION_DESC: true,
}

var SUPPORTED_QUERY_TYPE = map[int]bool{
	QUERY_TYPE_NAME:         true,
	QUERY_TYPE_MONSTER_TYPE: true,
}

var MAP_SORT_DIRECTION = map[int]string{
	SORT_DIRECTION_ASC:  "ASC",
	SORT_DIRECTION_DESC: "DESC",
}

var MAP_QUERY_TYPE = map[int]string{
	QUERY_TYPE_NAME:         "name",
	QUERY_TYPE_MONSTER_TYPE: "type",
}
