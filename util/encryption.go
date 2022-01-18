package util

import (
	"crypto/md5"
	"encoding/hex"
)

func GetMd5(s string, salt string) string {
	h := md5.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum([]byte(salt)))
}
