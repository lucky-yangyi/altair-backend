package service

import (
	"altair-backend/internal/dao"
	"altair-backend/internal/request"
	"altair-backend/internal/response"
	"altair-backend/pkg/utils"
	"context"
)

var GlobalCoinService *CoinService

//初始化
func GetCoinService() *CoinService {
	GlobalCoinService = newConService(dao.GetDao())
	return GlobalCoinService
}

type CoinService struct {
	baseService
}

func newConService(dao *dao.Dao) *CoinService {
	return &CoinService{baseService{dao: dao, ctx: context.Background()}}
}

//获取币种列表
func (s *CoinService) GetCoinList(ctx context.Context, req request.CoinListReq) (resp *response.CoinListResp, err error) {

	list, err := s.dao.GetCoinList(req.Name, req.Code, utils.GetPageOffset(req.PageNo, req.PageSize), req.PageSize)
	if err != nil {
		return nil, err
	}
	// 统计总数
	count, err := s.dao.CountCoinList(req.Name, req.Code)
	if err != nil {
		return nil, err
	}
	resp = &response.CoinListResp{
		List: list,
		Page: response.Page{
			PageNo:     req.PageNo,
			PageSize:   req.PageSize,
			TotalCount: uint64(count),
		},
	}
	return resp, nil
}

//增加币种列表
func (s *CoinService) CreateCoin(ctx context.Context, req request.CoinCreate) error {
	err := s.dao.CreateCoin(&req)
	if err != nil {
		return err
	}
	return nil
}
