// https://adventofcode.com/2015/day/4
package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"strings"
)

func lowest(input, prefix string) int {
	i := 0
	for {
		final := fmt.Sprintf("%s%d", input, i)
		hash := md5.New()
		io.WriteString(hash, final)
		result := hash.Sum(nil)
		str := hex.EncodeToString(result)
		if strings.HasPrefix(str, prefix) {
			break
		}
		i++
	}
	return i
}

func main() {
	fmt.Println(lowest("ckczppom", "00000"))
	fmt.Println(lowest("ckczppom", "000000"))
}
