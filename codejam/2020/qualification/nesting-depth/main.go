// Code Jam 2020
// Qualification Round
// Problem - Nesting Depth
// https://codingcompetitions.withgoogle.com/codejam/round/000000000019fd27/0000000000209a9f

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
		line := strings.Split(text[0:len(text)-1], "")
		arr := make([]int, 0)
		depth := 0
		var s bytes.Buffer
		
		for _, char := range line {
			num, _ := strconv.Atoi(char)
			arr = append(arr, num)
			for depth < num {
				s.WriteString("(")
				depth++
			}
			for depth > num {
				s.WriteString(")")
				depth--
			}
			s.WriteString(char)
		}

		for depth > 0 {
			s.WriteString(")")
			depth--
		}

		output.WriteString(fmt.Sprintf("Case #%d: %s\n", i, s.String()))
	}

	fmt.Print(output.String())

}
