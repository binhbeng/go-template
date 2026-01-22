package model

import (
	"github.com/binhbeng/goex/internal/dto"
	"github.com/binhbeng/goex/internal/global"
	"github.com/binhbeng/goex/internal/utils"
	"gorm.io/gorm"
	"gorm.io/plugin/soft_delete"
)

type BaseModel struct {
	ID        int64            `gorm:"column:id;type:int(11) unsigned AUTO_INCREMENT;not null;primarykey" json:"id"`
	CreatedAt utils.FormatDate `gorm:"column:created_at;type:timestamp;<-:create" json:"created_at"`
	UpdatedAt utils.FormatDate `gorm:"column:updated_at;type:timestamp" json:"updated_at"`
}

type BaseModelWithSoftDelete struct {
	BaseModel
	DeletedAt soft_delete.DeletedAt `gorm:"column:deleted_at;type:int(11) unsigned;not null;default:0;index;" json:"-"`
}

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		db: db,
	}
}

func (m *Repository) DB(model ...any) *gorm.DB {
	if model != nil {
		return m.db.Model(model[0])
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

func (m *Repository) Count(model any, condition string, args []any) (count int64, err error) {
	query := m.DB(model)
	if condition != "" {
		query = query.Where(condition, args...)
	}
	err = query.Count(&count).Error
	if err != nil {
		return 0, err
	}
	return
}
