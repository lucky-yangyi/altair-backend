package model

// MswStat 多签钱包每日统计
type MswStat struct {
	Mixin
	Wid       int     `gorm:"column:wid" json:"wid"`              // 多签id
	Date      string  `gorm:"column:date" json:"date"`            // 日期
	TransNum  int     `gorm:"column:trans_num" json:"transNum"`   // 交易数
	InNum     int     `gorm:"column:in_num" json:"inNum"`         // 收入笔数
	OutNum    int     `gorm:"column:out_num" json:"outNum"`       // 支出笔数
	InAmount  float64 `gorm:"column:in_amount" json:"inAmount"`   // 流入金额
	OutAmount float64 `gorm:"column:out_amount" json:"outAmount"` // 流出金额
}

// TableName get sql table name.获取数据库表名
func (m *MswStat) TableName() string {
	return "msw_stat"
}
