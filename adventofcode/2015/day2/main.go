// https://adventofcode.com/2015/day/2
package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var paper, ribbon int64
	for scanner.Scan() {
		text := strings.Split(scanner.Text(), "x")
		l, _ := strconv.Atoi(text[0])
		w, _ := strconv.Atoi(text[1])
		h, _ := strconv.Atoi(text[2])
		dims := []int{l, w, h}
		sort.Ints(dims)

		lw := l * w
		lh := l * h
		wh := w * h
		area := int64(2 * (lw + lh + wh))
		min := int64(math.Min(math.Min(float64(lw), float64(lh)), float64(wh)))
		paper += (area + min)

		volume := int64(l * w * h)
		distance := int64(2 * (dims[0] + dims[1]))
		ribbon += (volume + distance)
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	fmt.Println(paper)
	fmt.Println(ribbon)
}
