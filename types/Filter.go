package types

type FilterFormat struct {
	Query  map[string]string
	SortBy []string
	Order  []string
	Limit  int64
	Offset int64
}
