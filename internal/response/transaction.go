package response

import "altair-backend/internal/model"

type TransactionList struct {
	List []*model.MixTransaction
	Page
}

//type TransactionListDetail struct {
//	List  []*model.MixTransactionDetail
//}
