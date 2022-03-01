package response

type CollectNew struct {
	Address string `json:"address"` //钱包地址
}

type CollectAddress struct {
	Address []string `json:"address"` //钱包地址
}

type CollectTest struct {
	Message string `json:"Message"` //钱包地址
	Key     string `json:"Key"`     //钱包地址
	Sign    string `json:"Sign"`    //钱包地址
}
