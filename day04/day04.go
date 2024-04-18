package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strconv"
)

func main() {
	for i := 0; i < 10_000_000; i++ {
		text := "iwrupvqb" + strconv.Itoa(i)
		hash := GetMD5Hash(text)

		slice := hash[0:6]
		num, err := strconv.Atoi(slice)
		if err == nil {
			if num == 0 {
				fmt.Println(i)
			}
		}
	}
}

func GetMD5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}
