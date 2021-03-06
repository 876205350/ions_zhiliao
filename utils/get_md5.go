package utils

import (
	"crypto/md5"
	"fmt"
)

func GetMD5Str(str string) string {
	str_to_byte := []byte(str)
	byte_ret := md5.Sum(str_to_byte)
	ret := fmt.Sprintf("%x",byte_ret)
	return ret
}
