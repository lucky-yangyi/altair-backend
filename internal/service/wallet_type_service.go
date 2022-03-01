package service

import (
	"altair-backend/internal/dao"
	"altair-backend/internal/request"
	"altair-backend/internal/response"
	"altair-backend/pkg/utils"
	"context"
)

var GlobalWalletTypeService *WalletTypeService

//初始化
func GetWalletTypeService() *WalletTypeService {
	GlobalWalletTypeService = newWalletTypeService(dao.GetDao())
	return GlobalWalletTypeService
}

type WalletTypeService struct {
	baseService
}

func newWalletTypeService(dao *dao.Dao) *WalletTypeService {
	return &WalletTypeService{baseService{dao: dao, ctx: context.Background()}}
}

//获取币种列表
func (s *WalletTypeService) GetWalletTypeList(ctx context.Context, req request.WalletTypeListReq) (resp *response.WalletTypeListResp, err error) {

	list, err := s.dao.GetWalletTypeList(req.Name, req.Code, utils.GetPageOffset(req.PageNo, req.PageSize), req.PageSize)
	if err != nil {
		return nil, err
	}
	// 统计总数
	count, err := s.dao.CountWalletTypeList(req.Name, req.Code)
	if err != nil {
		return nil, err
	}
	resp = &response.WalletTypeListResp{
		List: list,
		Page: response.Page{
			PageNo:     req.PageNo,
			PageSize:   req.PageSize,
			TotalCount: uint64(count),
		},
	}
	return resp, nil
}
