// https://adventofcode.com/2015/day/3
package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	bytes, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	input := string(bytes)

	grid := make(map[string]bool)
	var x, y int
	var a, b int
	grid["0_0"] = true

	for i, c := range input {
		if i%2 == 0 {
			switch c {
			case '>':
				x++
				break
			case '^':
				y--
				break
			case 'v':
				y++
				break
			case '<':
				x--
				break
			}
			key := fmt.Sprintf("%d_%d", x, y)
			if _, ok := grid[key]; !ok {
				grid[key] = true
			}
		} else {
			switch c {
			case '>':
				a++
				break
			case '^':
				b--
				break
			case 'v':
				b++
				break
			case '<':
				a--
				break
			}
			key := fmt.Sprintf("%d_%d", a, b)
			if _, ok := grid[key]; !ok {
				grid[key] = true
			}
		}
	}

	fmt.Println(len(grid))
}
