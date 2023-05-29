package pagination

type PaginationParams struct {
	Page  int
	Limit int
}

func NewPaginationParams(page int, limit int) *PaginationParams {
	return &PaginationParams{Page: page, Limit: limit}
}
