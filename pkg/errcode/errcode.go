package errcode

import (
	"fmt"
)

type Error struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

var codes = map[int]string{}

func NewError(code int, msg string) *Error {
	if _, ok := codes[code]; ok {
		panic(fmt.Sprintf("错误码 %d 已经存在,请更换一个", code))
	}
	codes[code] = msg
	return &Error{Code: code, Msg: msg}

}
func (e Error) Error() string {
	return fmt.Sprintf("错误码: %d,错误码信息:%s", e.Code, e.Msg)

}
