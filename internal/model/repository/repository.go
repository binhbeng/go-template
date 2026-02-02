package repository

import (
	"github.com/binhbeng/goex/internal/dto"
	"github.com/binhbeng/goex/internal/global"
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		db: db,
	}
}

func (m *Repository) DB(entity ...any) *gorm.DB {
	if entity != nil {
		return m.db.Model(entity[0])
	}
	return m.db
}

func (m *Repository) Paginate(opt dto.PageOptionsDto) func(db *gorm.DB) *gorm.DB {
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

func (m *Repository) Count(entity any, condition string, args []any) (count int64, err error) {
	query := m.DB(entity)
	if condition != "" {
		query = query.Where(condition, args...)
	}
	err = query.Count(&count).Error
	if err != nil {
		return 0, err
	}
	return
}
