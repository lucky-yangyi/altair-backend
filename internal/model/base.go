package model

import (
	"database/sql/driver"
	"time"

	uuid "github.com/satori/go.uuid"
)

type LocalTime string
type Mixin struct {
	ID        uint64    `gorm:"column:id;type:bigint unsigned;primaryKey;not null;autoIncrement" json:"id"` //id
	CreatedAt LocalTime `gorm:"column:created_at;type:uint64;autoCreateTime;not null" json:"createdAt"`     // 创建时间
	UpdatedAt LocalTime `gorm:"column:updated_at;type:uint64;autoUpdateTime;not null" json:"updatedAt"`     // 修改时间
	IsDel     uint8     `gorm:"column:is_del;type:tinyint unsigned;not null;default:0" json:"isDel"`        // 0: 未删除 1: 删除
}

// Value 写入数据库之前，对数据做类型转换
func (s LocalTime) Value() (driver.Value, error) {
	loc := time.Local //设置时区
	if s == LocalTime("") {
		s = LocalTime(time.Now().Format("2006-01-02 15:04:05"))
	}
	tt, _ := time.ParseInLocation("2006-01-02 15:04:05", string(s), loc) //2006-01-02 15:04:05是转换的格式如php的"Y-m-d H:i:s"
	//fmt.Println(tt.Unix())
	return tt.Unix(), nil
}

// Scan 将数据库中取出的数据，赋值给目标类型
func (s *LocalTime) Scan(v interface{}) error {
	tt, ok := v.(time.Time)
	if ok {
		*s = LocalTime(tt.Format("2006-01-02 15:04:05"))
	} else {
		if v.(int64) > 315504000 {
			tTime := time.Unix(v.(int64), 0)
			*s = LocalTime(tTime.Format("2006-01-02 15:04:05"))
		} else {
			tTime := time.Unix(315504000, 0)
			*s = LocalTime(tTime.Format("2006-01-02 15:04:05"))
		}
	}
	return nil
}

type Uuid string

func (s Uuid) Value() (driver.Value, error) {
	return uuid.NewV4().String(), nil
}
