package cryptoutil

import (
	"crypto/md5"
	"encoding/hex"
	"io"
)

func MD5Str(s string) string {
	h := md5.New()
	io.WriteString(h, s)
	return hex.EncodeToString(h.Sum(nil))
}
