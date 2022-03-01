package response

// 公共分页
type Page struct {
	PageNo     uint64 `json:"pageNo"`     //页码
	PageSize   uint64 `json:"pageSize"`   //每页数量
	TotalCount uint64 `json:"totalCount"` //总数
}

type KeyInfo struct {
	Type       string `json:"type"`
	PrivateKey []byte `json:"privateKey"`
}
