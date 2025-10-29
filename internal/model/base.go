package model

import (
	"github.com/binhbeng/goex/data"
	"github.com/binhbeng/goex/internal/global"
	"github.com/binhbeng/goex/internal/pkg/utils"
	"gorm.io/gorm"
	"gorm.io/plugin/soft_delete"
)

type BaseModel struct {
	ID        uint             `gorm:"column:id;type:int(11) unsigned AUTO_INCREMENT;not null;primarykey" json:"id"`
	CreatedAt utils.FormatDate `gorm:"column:created_at;type:timestamp;<-:create" json:"created_at"`
	UpdatedAt utils.FormatDate `gorm:"column:updated_at;type:timestamp" json:"updated_at"`
}

func (m *BaseModel) DB(model ...any) *gorm.DB {
	return DB(model...)
}

func (m *BaseModel) Paginate(page, pageSize int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		offset := 0
		limit := global.PerPage
		if page < 1 {
			offset = page - 1
		}
		if pageSize > 0 {
			limit = pageSize
		}

		return db.Offset(offset * limit).Limit(limit)
	}
}

func (m *BaseModel) Count(model any, condition string, args []any) (count int64, err error) {
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

type ContainsDeleteBaseModel struct {
	BaseModel
	DeletedAt soft_delete.DeletedAt `gorm:"column:deleted_at;type:int(11) unsigned;not null;default:0;index;" json:"-"`
}

func DB(model ...any) *gorm.DB {
	if model != nil {
		return data.PostgreDB.Model(model[0])
	}
	return data.PostgreDB
}
