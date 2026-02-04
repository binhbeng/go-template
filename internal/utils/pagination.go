package utils

import (
	"github.com/binhbeng/goex/data"
	"github.com/binhbeng/goex/internal/dto"
	"github.com/binhbeng/goex/internal/global"
	"gorm.io/gorm"
)

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

func DB() *gorm.DB {
	return data.PostgreDB
}

func Paginate(opt dto.PageOptionsDto) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		// limit
		limit := global.PerPage
		if opt.Limit > 0 {
			limit = min(opt.Limit, global.PerPage)
		}

		// offset
		offset := 0
		if opt.Page > 1 {
			offset = (opt.Page - 1) * limit
		}

		// order
		if opt.OrderBy != "" {
			order := opt.OrderBy
			if opt.Direction != "" {
				order += " " + opt.Direction
			}
			db = db.Order(order)
		}

		return db.Offset(offset).Limit(limit)
	}
}
