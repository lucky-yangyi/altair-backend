package model

// CompanyMember 成员管理
type CompanyMember struct {
	Mixin
	CompanyID uint64 `gorm:"column:company_id;type:int;not null" json:"companyId"`           // 企业表ID
	Name      string `gorm:"column:name;type:varchar(255);default:''" json:"name"`           // 姓名
	Password  string `gorm:"column:password;type:varchar(50);default:''" json:"password"`    // 密码
	Email     string `gorm:"column:email;type:varchar(150);default:''" json:"email"`         // 邮箱
	Desc      string `gorm:"column:desc;type:varchar(255);default:''" json:"desc"`           // 备注
	IsAdmin   bool   `gorm:"column:is_admin;type:tinyint unsigned;default:0" json:"isAdmin"` // 1 管理员 0 普通
	Enabled   uint8  `gorm:"column:enabled;type:tinyint unsigned;default:1" json:"enabled"`  // 状态 1正常、0禁用
	IsDel     uint8  `gorm:"column:is_del;type:tinyint unsigned;default:0" json:"isDel"`     // 0: 未删除 1: 删除

	Company *Company `gorm:"foreignKey:CompanyID" json:"company"`
}

// CompanyMember 成员管理
type CompanyMemberNoPassword struct {
	Mixin
	CompanyID uint64 `gorm:"column:company_id;type:int;not null" json:"companyId"`           // 企业表ID
	Name      string `gorm:"column:name;type:varchar(255);default:''" json:"name"`           // 姓名
	Password  string `gorm:"column:password;type:varchar(50);default:''" json:"-"`           // 密码
	Email     string `gorm:"column:email;type:varchar(150);default:''" json:"email"`         // 邮箱
	Desc      string `gorm:"column:desc;type:varchar(255);default:''" json:"desc"`           // 备注
	IsAdmin   bool   `gorm:"column:is_admin;type:tinyint unsigned;default:0" json:"isAdmin"` // 1 管理员 0 普通
	Enabled   uint8  `gorm:"column:enabled;type:tinyint unsigned;default:1" json:"enabled"`  // 状态 1正常、0禁用
	IsDel     uint8  `gorm:"column:is_del;type:tinyint unsigned;default:0" json:"isDel"`     // 0: 未删除 1: 删除

	Company *Company `gorm:"foreignKey:CompanyID" json:"company"`
}

func (c CompanyMember) TableName() string {
	return "company_member"
}

// CMMeta no password
type CMMeta struct {
	ID   uint64 `gorm:"column:id;type:bigint unsigned;primaryKey;not null;autoIncrement" json:"id"` //id
	Name string `gorm:"column:name;type:varchar(255);default:''" json:"name"`                       // 姓名
}

func (c CMMeta) TableName() string {
	return "company_member"
}
