package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

func Bitcoin(leadingZeros int, key string) int {
	counter := 0
	for {
		counter++
		compoundKey := fmt.Sprintf("%s%d", key, counter)
		hasher := md5.New()
		hasher.Write([]byte(compoundKey))
		hashed := hex.EncodeToString(hasher.Sum(nil))
		leadingString := ""
		for i := 0; i < leadingZeros; i++ {
			leadingString += "0"
		}
		if hashed[0:leadingZeros] == leadingString {
			return counter
		}
	}
}
