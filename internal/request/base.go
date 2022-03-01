package request

type PageReq struct {
	PageNo   uint64 `json:"pageNo" binding:"required,gt=0"`   //页码数
	PageSize uint64 `json:"pageSize" binding:"required,gt=4"` //每页数量
}
type Ids struct {
	Id []int `json:"id" binding:"required"`
}
