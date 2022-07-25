package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func MD5(v string) string {
	d := []byte(v)
	m := md5.New()
	m.Write(d)
	return hex.EncodeToString(m.Sum(nil))
}

func main() {
	pwd := "lcg9899"
	bytes, _ := bcrypt.GenerateFromPassword([]byte(pwd), 14)
	fmt.Println(string(bytes))
}
