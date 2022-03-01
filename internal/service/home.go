package service

import (
	"altair-backend/internal/dao"
	"altair-backend/internal/model"
	"altair-backend/internal/response"
	"altair-backend/log"
	"time"
)

func TotalStat(msw []model.Wallet) (ids []uint64, TotalStat response.TotalStat) {
	for _, v := range msw {
		TotalStat.TotalAmount += v.Balance
		ids = append(ids, v.ID)
	}

	// TotalStat.WalletNum
	TotalStat.WalletNum = len(ids)

	// TotalStat.ThreeMonthOut/ThreeMonthIn 三个月收入和支出
	var threeStat []model.MswStat
	tm := time.Now()
	threeMonthDate := tm.AddDate(0, -3, 0).Format("2006-01-02")
	today := tm.Format("2006-01-02")
	dao.DB.Model(model.MswStat{}).
		Where("wid IN ?", ids).
		Where("date >= ?", threeMonthDate).
		Where("date <= ?", today).
		Find(&threeStat)
	for _, v := range threeStat {
		TotalStat.ThreeMonthOut += v.OutAmount
		TotalStat.ThreeMonthIn += v.InAmount
	}
	return
}

// 3
func MultiSignWalletList(msw []model.Wallet, totalAmount float64) (res []response.Msw, pipe []response.Pipe) {

	for _, v := range msw {
		log.Info("address:", v.Address)
		log.Info("Balance:", v.Balance)
		var singleMsw response.Msw
		var p response.Pipe
		singleMsw.ID = v.ID
		singleMsw.Balance = v.Balance
		singleMsw.Address = v.Address
		singleMsw.Name = v.Name
		p.Value = v.Balance
		p.Name = v.Name

		if totalAmount == 0 {
			singleMsw.Percent = 0
		} else {
			singleMsw.Percent = v.Balance / totalAmount
		}

		//3个月的总额
		var threeStat []model.MswStat
		tm := time.Now()
		// 从今天往前推3个月
		threeMonthDate := tm.AddDate(0, -3, 0).Format("2006-01-02")
		today := tm.Format("2006-01-02")
		dao.DB.Model(model.MswStat{}).
			Where("wid = ?", v.ID).
			Where("date >= ?", threeMonthDate).
			Where("date <= ?", today).
			Find(&threeStat)
		for _, v := range threeStat {
			singleMsw.ThreeMonthOut += v.OutAmount
			singleMsw.ThreeMonthIn += v.InAmount
			singleMsw.OutNum += v.OutNum
			singleMsw.InNum += v.InNum
		}
		res = append(res, singleMsw)
		pipe = append(pipe, p)
	}
	return
}

func SixMonthAmountList(ids []uint64) (list []response.SixMonthAmount) {
	for i := 5; i >= 0; i-- {
		var sma response.SixMonthAmount
		var mswStat []model.MswStat
		iMonthBegin, iMonthEnd, iMonth := GetMonthBeginAndEnd(i)
		sma.Month = iMonth
		dao.DB.Model(model.MswStat{}).
			Where("wid IN ?", ids).
			Where("date >= ?", iMonthBegin).
			Where("date <= ?", iMonthEnd).
			Find(&mswStat)
		for _, v := range mswStat {
			sma.In += v.InAmount
			sma.Out += v.OutAmount
		}
		list = append(list, sma)
	}
	return
}

func ThreeMonthNum(Msw []response.Msw) (data response.ThreeMonthNum) {

	for _, v := range Msw {
		data.In += uint64(v.InNum)
		data.Out += uint64(v.OutNum)
	}
	return
}

func FifteenDayAmountInOutList(ids []uint64) (list []response.FifteenDayInOutNum) {
	for i := 14; i >= 0; i-- {
		var sma response.FifteenDayInOutNum
		var mswStat model.MswStat
		iDate := GetDate(i)
		sma.Date = iDate
		for _, key := range ids {
			dao.DB.Model(model.MswStat{}).
				Where("wid = ?", key).
				Where("date = ?", iDate).
				Find(&mswStat)
			var detail response.WalletDetail
			detail.ID = key
			detail.Out = mswStat.OutAmount
			detail.In = mswStat.InAmount
			detail.OutNum = mswStat.OutNum
			detail.InNum = mswStat.InNum
			sma.WalletDetail = append(sma.WalletDetail, detail)
		}

		list = append(list, sma)
	}
	return
}

func FifteenDayAmountInOutListById(ids []uint64) (list []response.FifteenDayInOutNumById) {
	for i := 14; i >= 0; i-- {
		var mswStat []model.MswStat
		iDate := GetDate(i)
		var detail response.FifteenDayInOutNumById
		detail.Date = iDate
		dao.DB.Model(model.MswStat{}).
			Where("mid IN ?", ids).
			Where("date = ?", iDate).
			Find(&mswStat)
		for _, v := range mswStat {
			detail.Out += v.OutAmount
			detail.In += v.InAmount
			detail.OutNum += v.OutNum
			detail.InNum += v.InNum
		}

		list = append(list, detail)
	}
	return
}

// GetMonthBeginAndEnd get the day head and tail
func GetMonthBeginAndEnd(monthBefore int) (string, string, string) {
	now := time.Now()
	currentYear, currentMonth, _ := now.AddDate(0, -monthBefore, 0).Date()
	currentLocation := now.Location()
	firstOfMonth := time.Date(currentYear, currentMonth, 1, 0, 0, 0, 0, currentLocation)
	lastOfMonth := firstOfMonth.AddDate(0, 1, -1)
	iMonthBegin := firstOfMonth.Format("2006-01-02")
	iMonthEnd := lastOfMonth.Format("2006-01-02")
	iMonth := firstOfMonth.Format("2006-01")
	return iMonthBegin, iMonthEnd, iMonth
}

func GetDate(dayBefore int) (day string) {
	now := time.Now()
	day = now.AddDate(0, 0, -dayBefore).Format("2006-01-02")
	return
}
