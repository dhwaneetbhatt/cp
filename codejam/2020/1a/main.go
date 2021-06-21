package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var text string

	text, _ = reader.ReadString('\n')
	t, _ := strconv.Atoi(text[0 : len(text)-1])
	var output bytes.Buffer

	for i := 1; i <= t; i++ {
		text, _ = reader.ReadString('\n')
		n, _ := strconv.Atoi(text[0 : len(text)-1])
		arr := make([]string, t)

		longest := ""
		for j := 0; j < n; j++ {
			text, _ = reader.ReadString('\n')
			text = text[0 : len(text)-1]
			arr = append(arr, text)

			sp := strings.Replace(text, "*", "", -1)
			if len(sp) > len(longest) {
				longest = sp
			}
		}

		s := longest
		for _, text := range arr {
			split := strings.Split(strings.TrimSpace(text), "*")
			if strings.TrimSpace(split[0]) != "" {
				if !strings.HasPrefix(longest, split[0]) {
					s = "*"
					break
				}
			}
			if len(split) > 1 && strings.TrimSpace(split[1]) != "" {
				if !strings.HasSuffix(longest, split[1]) {
					s = "*"
					break
				}
			}
		}

		output.WriteString(fmt.Sprintf("Case #%d: %s\n", i, strings.TrimSpace(s)))
	}

	fmt.Print(output.String())

}
