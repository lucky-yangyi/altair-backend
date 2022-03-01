package model

// Admin 运营后台
type Admin struct {
	Mixin
	Name     string `gorm:"column:name;type:varchar(255);default:''" json:"name"`           // 用户名
	Email    string `gorm:"column:email;type:varchar(150);default:''" json:"email"`         // 邮箱
	Password string `gorm:"column:password;type:varchar(50);default:''" json:"password"`    // 密码
	IsAdmin  uint8  `gorm:"column:is_admin;type:tinyint unsigned;default:0" json:"isAdmin"` // 0 超级管理员 1 普通管理员
	Enabled  uint8  `gorm:"column:enabled;type:tinyint unsigned;default:0" json:"enabled"`  // 状态 0正常、1禁用
}

type AdminNoPassword struct {
	Mixin
	Name     string `gorm:"column:name;type:varchar(255);default:''" json:"name"`           // 用户名
	Email    string `gorm:"column:email;type:varchar(150);default:''" json:"email"`         // 邮箱
	Password string `gorm:"column:password;type:varchar(50);default:''" json:"-"`           // 密码
	IsAdmin  uint8  `gorm:"column:is_admin;type:tinyint unsigned;default:0" json:"isAdmin"` // 0 管理员 1 普通
	Enabled  uint8  `gorm:"column:enabled;type:tinyint unsigned;default:0" json:"enabled"`  // 状态 0正常、1禁用
}
