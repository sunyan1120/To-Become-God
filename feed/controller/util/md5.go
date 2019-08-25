package util

import (
	"crypto/md5"
	"encoding/hex"
)

//使用md5对密码进行加密
func Encryption(pwd string) string {
	h := md5.New()
	h.Write([]byte(pwd))
	cipherStr := h.Sum(nil)
	return hex.EncodeToString(cipherStr)
}
