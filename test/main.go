package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strconv"
	"time"

	"liandyuan.cn/api/util"
)

func MD5(v string) string {
	d := []byte(v)
	m := md5.New()
	m.Write(d)
	return hex.EncodeToString(m.Sum(nil))
}

func main() {
	tm := time.Now().Unix()
	times := strconv.FormatInt(tm, 10)
	userid := util.MD5(times)
	user := util.MD5(string(tm))
	fmt.Println(userid)
	fmt.Println(times)
	fmt.Println(user)

}
