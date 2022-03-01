package response

type TotalStat struct {
	TotalAmount   float64
	ThreeMonthIn  float64
	ThreeMonthOut float64
	WalletNum     int
}

// 2 6个月的收入支出金额

type SixMonthAmount struct {
	Month string
	In    float64
	Out   float64
}

type InOut struct {
	In  float64
	Out float64
}

// 2 3

type Msw struct {
	ID            uint64
	Name          string
	Address       string
	Balance       float64
	ThreeMonthIn  float64
	ThreeMonthOut float64
	Percent       float64
	InNum         int
	OutNum        int
}

type Pipe struct {
	Value float64 `json:"value"`
	Name  string  `json:"name"`
}

// 4
//chain transaction list

// 5 三个月的收入支出次数

type ThreeMonthNum struct {
	In  uint64
	Out uint64
}

// 6 15天的收入支出次数详情

type FifteenDayInOutNum struct {
	Date         string
	WalletDetail []WalletDetail
}

type FifteenDayInOutNumById struct {
	Date   string
	In     float64
	Out    float64
	InNum  int
	OutNum int
}

type WalletDetail struct {
	ID     uint64
	In     float64
	Out    float64
	InNum  int
	OutNum int
}

type Stat struct {
	TotalStat          TotalStat
	SixMonthAmount     []SixMonthAmount
	Msw                []Msw
	ThreeMonthNum      ThreeMonthNum
	FifteenDayInOutNum []FifteenDayInOutNum
	Pipe               []Pipe
}

type StatById struct {
	TotalStat          TotalStat
	SixMonthAmount     []SixMonthAmount
	Msw                []Msw
	ThreeMonthNum      ThreeMonthNum
	FifteenDayInOutNum []FifteenDayInOutNumById
	Pipe               []Pipe
}
