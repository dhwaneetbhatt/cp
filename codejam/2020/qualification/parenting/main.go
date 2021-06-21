// Code Jam 2020
// Qualification Round
// Problem - Parenting Partnering Returns
// https://codingcompetitions.withgoogle.com/codejam/round/000000000019fd27/000000000020bdf9

package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type interval struct {
	idx       int
	startTime int
	endTime   int
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	var text string

	text, _ = reader.ReadString('\n')
	t, _ := strconv.Atoi(text[0 : len(text)-1])
	var output bytes.Buffer

	for i := 1; i <= t; i++ {
		text, _ = reader.ReadString('\n')
		n, _ := strconv.Atoi(text[0 : len(text)-1])
		intervals := make([]interval, n)
		assign := make([]string, n)
		possible := true

		for j := 0; j < n; j++ {
			text, _ = reader.ReadString('\n')
			line := strings.Split(text[0:len(text)-1], " ")
			start, _ := strconv.Atoi(line[0])
			end, _ := strconv.Atoi(line[1])
			intervals[j] = interval{
				idx:       j,
				startTime: start,
				endTime:   end,
			}
		}

		intervals = sort(intervals)

		var cEnd, jEnd int = 0, 0
		for _, i := range intervals {
			if i.startTime >= cEnd {
				assign[i.idx] = "C"
				cEnd = i.endTime
			} else if i.startTime >= jEnd {
				assign[i.idx] = "J"
				jEnd = i.endTime
			} else {
				possible = false
				break
			}
		}

		result := ""
		if possible {
			result = strings.Join(assign, "")
		} else {
			result = "IMPOSSIBLE"
		}
		output.WriteString(fmt.Sprintf("Case #%d: %s\n", i, result))
	}

	fmt.Print(output.String())

}

func sort(a []interval) []interval {
	if len(a) < 2 {
		return a
	}

	left, right := 0, len(a)-1
	pivot := 0

	a[pivot], a[right] = a[right], a[pivot]

	for i := range a {
		if a[i].startTime < a[right].startTime {
			a[left], a[i] = a[i], a[left]
			left++
		}
	}

	a[left], a[right] = a[right], a[left]

	sort(a[:left])
	sort(a[left+1:])

	return a
}
