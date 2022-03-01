package task

import (
	"altair-backend/internal/dao"
	"altair-backend/internal/model"
	"fmt"
	"math"
	"strconv"
	"time"
)

const SUCCESS = 1 //成功

func MswStat() error {
	fmt.Println("MswStat: start save")
	// 获取msw列表
	WalletSlice := GetMswListSlice()

	// date string
	now := time.Now()
	date := now.AddDate(0, 0, -1)

	dateStr := date.Format("2006-01-02")

	tBegin, tEnd := GetDayTimestamp(dateStr)

	// mysql work
	for _, v := range WalletSlice {
		SaveAddressToAndFrom(int(v.ID), dateStr, v.Address, tBegin, tEnd)
	}

	// 获取时间区间

	return nil
}

func SaveAddressToAndFrom(mid int, date string, address string, tBegin, tEnd int64) {
	// 支出
	var chainFrom []model.ChainTransaction
	dao.DB.Model(model.ChainTransaction{}).
		Where("from_address = ?", address).
		Where("timestamp > ?", tBegin).
		Where("timestamp < ?", tEnd).
		Where("status = ?", SUCCESS).
		Find(&chainFrom)

	var stat model.MswStat
	stat.Wid = mid
	stat.Date = date

	for _, vo := range chainFrom {
		value, _ := strconv.ParseFloat(vo.Value, 64)
		stat.OutAmount += value / math.Pow(10, 18)
		stat.OutNum++
	}

	// 收入
	var chainTo []model.ChainTransaction
	dao.DB.Model(model.ChainTransaction{}).
		Where("to_address = ?", address).
		Where("timestamp > ?", tBegin).
		Where("timestamp < ?", tEnd).
		Where("status = ?", SUCCESS).
		Find(&chainTo)

	for _, vi := range chainTo {
		value, _ := strconv.ParseFloat(vi.Value, 64)
		stat.InAmount += value / math.Pow(10, 18)
		stat.InNum++
	}
	stat.TransNum = stat.InNum + stat.OutNum

	// save stat
	// update or create
	var ExitStat model.MswStat
	dao.DB.Model(model.MswStat{}).Where("mid = ?", stat.Wid).Where("date = ?", stat.Date).Find(&ExitStat)
	if ExitStat.ID != 0 {
		stat.ID = ExitStat.ID
		ExitStat = stat
		dao.DB.Save(&ExitStat)
	} else {
		dao.DB.Model(model.MswStat{}).Create(&stat)
	}

}

// GetDayTimestamp get date begin & end timestamp
func GetDayTimestamp(date string) (int64, int64) {
	// 字符串解析成时间格式
	t, _ := time.Parse("2006-01-02", date)
	tsBegin := t.Unix()

	tsEnd := tsBegin + 86400 - 1
	return tsBegin, tsEnd

}

func GetMswListSlice() (WalletList []model.Wallet) {
	dao.DB.Model(model.Wallet{}).Find(&WalletList)
	return
}
