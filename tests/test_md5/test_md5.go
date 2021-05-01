package main

import (
	"crypto/md5"
	"fmt"
)

func GetMD5File(str string) string {
	str_to_byte := []byte(str)
	byte_ret := md5.Sum(str_to_byte)
	ret := fmt.Sprintf("%x",byte_ret)
	return ret
}

func main()  {
	md5_ret := GetMD5File("12345678")
	fmt.Println(md5_ret)
}
