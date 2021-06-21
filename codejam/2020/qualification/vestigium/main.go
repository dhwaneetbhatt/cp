// Code Jam 2020
// Qualification Round
// Problem - Vestigium
// https://codingcompetitions.withgoogle.com/codejam/round/000000000019fd27/000000000020993c

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
		matrix := make([][]int, n)
		sum := 0
		rowDup := 0
		colDup := 0

		for j := 0; j < n; j++ {

			text, _ = reader.ReadString('\n')
			line := strings.Split(text[0:len(text)-1], " ")
			dup := make(map[int]bool)
			row := make([]int, n)
			isRowDup := false

			for k, char := range line {
				num, _ := strconv.Atoi(char)
				row[k] = num
				if j == k {
					sum += num
				}
				if _, ok := dup[num]; ok {
					isRowDup = true
				}
				dup[num] = true
			}

			matrix[j] = row

			if isRowDup {
				rowDup++
			}
		}

		for j := 0; j < n; j++ {

			dup := make(map[int]bool)
			isColDup := false

			for k := 0; k < n; k++ {
				num := matrix[k][j]
				if _, ok := dup[num]; ok {
					isColDup = true
				}
				dup[num] = true
			}

			if isColDup {
				colDup++
			}
		}

		output.WriteString(fmt.Sprintf("Case #%d: %d %d %d\n", i, sum, rowDup, colDup))
	}

	fmt.Print(output.String())
}
