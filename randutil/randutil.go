package randutil

import (
	"math/rand"
	"time"
)

var (
	alnum       = `0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz`
	defaultRand = rand.New(rand.NewSource(time.Now().UnixNano()))
)

func RandAlnumStr(n int) string {
	return string(RandAlnum(n))
}

func RandAlnum(n int) []byte {
	b := make([]byte, n)
	for i := 0; i < n; i++ {
		b[i] = alnum[defaultRand.Intn(len(alnum))]
	}
	return b
}
