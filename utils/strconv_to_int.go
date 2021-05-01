package utils

import (
	"strconv"
	"unsafe"
)

func StrToInt(str_name string) int {
	id_int64,_ := strconv.ParseInt(str_name,10,64)
	id_int := *(*int)(unsafe.Pointer(&id_int64))
	return  id_int
}
