package model

// Company 企业管理
type Company struct {
	Mixin
	Name    string `gorm:"column:name;type:varchar(50);default:''" json:"name"`           // 企业名称
	Enabled uint8  `gorm:"column:enabled;type:tinyint unsigned;default:1" json:"enabled"` // 状态 0禁用、1正常
}

// TableName get sql table name.获取数据库表名
func (m *Company) TableName() string {
	return "company"
}

//func (m *Company) GetCompanyById(id uint64) {
//	dao.DB.Table(m.TableName()).Where("id = ?", id).First(m)
//}