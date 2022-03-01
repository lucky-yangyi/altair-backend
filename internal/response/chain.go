package response

type Status struct {
	Code    int    `json:"code"`
	Success bool   `json:"success"`
	Error   string `json:"error"`
}

type CommonResponse struct {
	Status
	Data interface{} `json:"data"`
}

type CollectTransactionResponse struct {
	Status
	Data CollectTransaction `json:"data"`
}

type CollectTransaction struct {
	Cid    string  `json:"cid"`
	Amount float64 `json:"amount"`
}

type ReceiptResponse struct {
	Status
	Data Receipt `json:"data"`
}
type Receipt struct {
}
