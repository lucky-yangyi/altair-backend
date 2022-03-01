package utils

import (
	"fmt"
	"gorm.io/gorm"
	"time"
)

// Paginate preload的时候不能一次取出
func Paginate(queryTx *gorm.DB, page uint64, perPage uint64, model interface{}) (Page, *gorm.DB, error) {
	var total int64
	err := queryTx.Count(&total).Error

	queryTx = queryTx.Scopes(Pg(page, perPage))

	if err != nil {
		fmt.Println(err)
	}

	return Page{
		List:       model,
		PageNo:     page,
		PageSize:   perPage,
		TotalCount: total,
	}, queryTx, err
}

type Page struct {
	PageNo     uint64      `json:"pageNo"`
	PageSize   uint64      `json:"pageSize"`
	TotalCount int64       `json:"totalCount"`
	List       interface{} `json:"list"`
}

func Pg(page uint64, pageSize uint64) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		//默认page 1
		if page == 0 {
			page = 1
		}
		// 默认20条一页
		if pageSize == 0 {
			pageSize = 20
		}
		offset := (page - 1) * pageSize
		return db.Offset(int(offset)).Limit(int(pageSize))
	}
}

// StringToTimestamp 时间戳转换处理
func StringToTimestamp(s string) int64 {
	loc := time.Local //设置时区
	tt, _ := time.ParseInLocation("2006-01-02", s, loc)
	return tt.Unix()
}
