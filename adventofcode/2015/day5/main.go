// https://adventofcode.com/2015/day/5
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	nice := 0
	for scanner.Scan() {
		input := scanner.Text()
		if isNice1(input) {
			nice++
		}
	}
	fmt.Println(nice)
}
