package model

type Coin struct {
	Mixin
	Name string `json:"name"` //币种名
	Code string `json:"code"` //币种代号
}

// TableName get sql table name.获取数据库表名
func (m *Coin) TableName() string {
	return "coin"
}
