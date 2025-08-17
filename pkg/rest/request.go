package rest

import (
	"fmt"

	"gorm.io/gorm"
)

type RequestPagination struct {
	Page   int    `form:"page"`
	Limit  int    `form:"limit"`
	Search string `form:"search"`
}

func (r *RequestPagination) GetOffset() int {
	return (r.Page - 1) * r.Limit
}

func (req *RequestPagination) SliceItems(query *gorm.DB) *gorm.DB {
	return query.Limit(req.Limit).Offset(req.GetOffset())
}

func (req *RequestPagination) FilterSearch(query *gorm.DB, search_fields []string) *gorm.DB {
	if req.Search != "" && len(search_fields) > 0 {
		likePattern := "%" + req.Search + "%"
		for i, field := range search_fields {
			if i == 0 {
				query = query.Where(fmt.Sprintf("%s ILIKE ?", field), likePattern)
			} else {
				query = query.Or(fmt.Sprintf("%s ILIKE ?", field), likePattern)
			}
		}
	}
	return query
}
