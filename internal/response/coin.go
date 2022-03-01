package response

import "altair-backend/internal/model"

type CoinListResp struct {
	List []*model.Coin
	Page
}
