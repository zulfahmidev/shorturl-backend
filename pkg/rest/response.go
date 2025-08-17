package rest

import (
	"math"
)

type BaseResponse struct {
	Message string      `json:"message"`
	Body    interface{} `json:"body"`
}

type BaseListResponse struct {
	Items      interface{} `json:"items"`
	TotalPages int64       `json:"totalPages"`
	TotalItems int64       `json:"totalItems"`
	Page       int         `json:"page"`
	Limit      int         `json:"limit"`
}

type MetaResp struct {
	Page  int `json:"page"`
	Limit int `json:"limit"`
}

func BuildListResponse(items interface{}, totalItem int64, meta_resp MetaResp) BaseListResponse {
	totalPages := math.Ceil(float64(totalItem) / float64(meta_resp.Limit))
	return BaseListResponse{
		Items:      items,
		TotalPages: int64(totalPages),
		TotalItems: totalItem,
		Page:       meta_resp.Page,
		Limit:      meta_resp.Limit,
	}
}
