package dto

type PageOptionsDto struct {
	Page      int    `form:"page" json:"page" binding:"required"`
	Limit     int    `form:"limit" json:"limit" binding:"omitempty"`
	OrderBy   string `form:"orderBy" json:"orderBy"`
	Direction string `form:"direction" json:"direction" binding:"omitempty,oneof=asc ASC desc DESC" message:"direction must be asc or desc"`
}
