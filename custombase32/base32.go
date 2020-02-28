package custombase32

import "encoding/base32"

func Encoder() *base32.Encoding {
	return base32.NewEncoding("abcdefghijklmnopqrstuvwxyz234567")
}
