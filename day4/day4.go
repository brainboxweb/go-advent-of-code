package day4

import (
	"crypto/md5" // #nosec G501
	"encoding/hex"
	"fmt"
)

func Bitcoin(leadingZeros int, key string) int {
	counter := 0
	for {
		counter++
		compoundKey := fmt.Sprintf("%s%d", key, counter)
		hasher := md5.New() // #nosec G401
		_, err := hasher.Write([]byte(compoundKey))
		if err != nil {
			panic("not expected")
		}
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
