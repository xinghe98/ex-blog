package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

func MD5(v string) string {
	d := []byte(v)
	m := md5.New()
	m.Write(d)
	return hex.EncodeToString(m.Sum(nil))
}

func main() {

	a := MD5("1111111111111111")
	fmt.Println(a)
}
