package model

type PageReq struct {
	Page     int `json:"page"`
	PageSize int `json:"page_size"`
}
