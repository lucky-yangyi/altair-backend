package dao

import "gorm.io/gorm"

var GlobalDao *Dao
var DB *gorm.DB

func init() {
	DB = newMysql()
}

type Dao struct {
	dao *gorm.DB
}

func newDao(db *gorm.DB) *Dao {
	return &Dao{dao: db}
}

// 初始化
func GetDao() *Dao {
	GlobalDao = newDao(DB)
	return GlobalDao
}

func (d *Dao) CreateTx() *gorm.DB {
	return d.dao.Begin()
}

func (d *Dao) Rollback(db *gorm.DB) {
	if db == nil {
		return
	}
	db.Rollback()
}

func (d *Dao) Commit(db *gorm.DB) {
	if db == nil {
		return
	}
	db.Commit()
}
