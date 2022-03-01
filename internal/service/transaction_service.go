package service

import (
	"altair-backend/internal/dao"
	"altair-backend/internal/model"
	"altair-backend/internal/request"
	"altair-backend/internal/response"
	"altair-backend/pkg/utils"
	"context"
)

var GlobalTransactionService *TransactionService

//初始化
func GetTransactionService() *TransactionService {
	GlobalTransactionService = newTransactionService(dao.GetDao())
	return GlobalTransactionService
}

type TransactionService struct {
	baseService
}

func newTransactionService(dao *dao.Dao) *TransactionService {
	return &TransactionService{baseService{dao: dao, ctx: context.Background()}}
}

// 获取交易列表
func (s *TransactionService) GetTransaction(ctx context.Context, req request.TransactionList) (resp *response.TransactionList, err error) {
	list, err := s.dao.GetTransactionList(req, utils.GetPageOffset(req.PageNo, req.PageSize), req.PageSize)
	if err != nil {
		return nil, err
	}
	count, err := s.dao.CountTransaction(req)
	if err != nil {
		return nil, err
	}
	resp = &response.TransactionList{
		List: list,
		Page: response.Page{
			PageNo:     req.PageNo,
			PageSize:   req.PageSize,
			TotalCount: uint64(count),
		},
	}
	return
}

// 获取交易详情
func (s *TransactionService) GetTransactionDetail(ctx context.Context, req request.TransactionDetail) (resp *model.TransactionDetail, err error) {
	resp, err = s.dao.GetTransactionDetail(req.TransactionId)
	if err != nil {
		return nil, err
	}
	return
}
