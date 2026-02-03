package utils

type Pagination struct {
	Page         int  `json:"page"`
	Limit        int  `json:"limit"`
	TotalRecords int  `json:"total_records"`
	TotalPages   int  `json:"total_pages"`
	HasNext      bool `json:"has_next"`
	HasPrev      bool `json:"has_prev"`
}

func NewPagiantion(page, limit, totalRecords int) *Pagination {
	if page <= 0 {
		page = 1
	}

	totalPages := (totalRecords + limit - 1) / limit

	return &Pagination{
		Page:         page,
		Limit:        limit,
		TotalRecords: totalRecords,
		TotalPages:   totalPages,
		HasNext:      page < totalPages,
		HasPrev:      page > 1,
	}
}

func NewPagiantionResponse(data any, page, limit, totalRecords int) map[string]any {
	return map[string]any{
		"data":       data,
		"pagination": NewPagiantion(page, limit, totalRecords),
	}
}
