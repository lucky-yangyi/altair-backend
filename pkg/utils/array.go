package utils

func IsInArray(s interface{}, d []int) bool {
	for _, v := range d {
		if s == v {
			return true
		}
	}
	return false
}
