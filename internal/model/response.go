package model

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

type PageResp struct {
	Page     int
	PageSize int
	Count    int
}

type timestr struct {
	create string
	update string
}
