package service

import (
	"altair-backend/internal/dao"
	"altair-backend/internal/request"
	"altair-backend/internal/response"
	"altair-backend/pkg/utils"
	"context"
)

var GlobalBillService *BillService

//初始化
func GetBillService() *BillService {
	GlobalBillService = newBillService(dao.GetDao())
	return GlobalBillService
}

type BillService struct {
	baseService
}

func newBillService(dao *dao.Dao) *BillService {
	return &BillService{baseService{dao: dao, ctx: context.Background()}}
}

//运营后台月付账单管理
func (s *BillService) GetBillList(ctx context.Context, req request.BillListReq) (resp *response.MonthListResp, err error) {
	list, err := s.dao.GetBillList(req.WalletId, req.PayStatus, req.Month, utils.GetPageOffset(req.PageNo, req.PageSize), req.PageSize)
	if err != nil {
		return nil, err
	}
	count, err := s.dao.GetBillCount(req.WalletId, req.PayStatus, req.Month)
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

//账单缴费
func (s *BillService) StatusBill(ctx context.Context, req request.StatusBillReq) (err error) {
	return s.dao.UpdateStatusBill(1, req.Id)
}
