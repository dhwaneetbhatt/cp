// https://adventofcode.com/2015/day/1
package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	b, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	input := string(b)

	floor := 0
	position := 0
	for i, c := range input {
		if c == '(' {
			floor++
		} else if c == ')' {
			floor--
		}
		if floor == -1 && position == 0 {
			position = i + 1
		}
	}
	fmt.Println(floor)
	fmt.Println(position)
}
