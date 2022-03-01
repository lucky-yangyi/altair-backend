package service

import (
	"altair-backend/internal/dao"
	"altair-backend/internal/model"
	"altair-backend/internal/request"
	"altair-backend/internal/response"
	"altair-backend/pkg/utils"
	"context"
)

var GlobalMonthBillService *MonthBillService

//初始化
func GetMonthBillService() *MonthBillService {
	GlobalMonthBillService = newMonthBillService(dao.GetDao())
	return GlobalMonthBillService
}

type MonthBillService struct {
	baseService
}

func newMonthBillService(dao *dao.Dao) *MonthBillService {
	return &MonthBillService{baseService{dao: dao, ctx: context.Background()}}
}

//获取月付账单
func (s *MonthBillService) GetMonthBillList(ctx context.Context, req request.MonthBillListReq) (resp *response.MonthListResp, err error) {
	list, err := s.dao.GetMonthBillList(req.WalletId, req.PayStatus, req.Month, utils.GetPageOffset(req.PageNo, req.PageSize), req.PageSize)
	if err != nil {
		return nil, err
	}
	count, err := s.dao.GetMonthBillCount(req.WalletId, req.PayStatus, req.Month)
	if err != nil {
		return nil, err
	}
	resp = &response.MonthListResp{
		List: list,
		Page: response.Page{
			PageNo:     req.PageNo,
			PageSize:   req.PageSize,
			TotalCount: uint64(count),
		},
	}
	return
}

func (s *MonthBillService) AddMonthBill() {
	var companyId uint64
	var err error
	var wallet *model.MonthBill
	companyId = 12
	list, err := s.dao.GetWalletList(companyId)
	if err != nil {
		return
	}
	for _, v := range list {
		wallet = &model.MonthBill{
			Name:      v.Name,
			Address:   v.Address,
			Month:     "",
			Amount:    0.01,
			PayStatus: 0,
			IsDel:     0,
		}
		err := s.dao.CreateMonthBill(wallet)
		if err != nil {
			return
		}
	}
}
