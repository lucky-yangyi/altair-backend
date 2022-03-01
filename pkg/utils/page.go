package utils

func GetPageOffset(page, pageSize uint64) uint64 {
	//页码数  1
	//每页数量 10
	var result uint64
	if page > 0 {
		result = (page - 1) * pageSize
	}
	return result
}

func CheckIsPage(page, size uint64) (uint64, uint64) {
	if page == 0 {
		page = 1
	}
	if size <= 0 || size > 100 {
		size = 10
	}
	return page, size
}
