package main

import "strings"

var badStrings = []string{"ab", "cd", "pq", "xy"}

func countVowels(input string) int {
	vowels := 0
	for _, c := range input {
		if c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u' {
			vowels++
		}
	}
	return vowels
}

func hasConsecutiveLetters(input string) bool {
	for i := 0; i < len(input)-1; i++ {
		if input[i] == input[i+1] {
			return true
		}
	}
	return false
}

func hasBadStrings(input string, banned []string) bool {
	for _, s := range banned {
		if strings.Contains(input, s) {
			return true
		}
	}
	return false
}

func isNice1(input string) bool {
	if !hasBadStrings(input, badStrings) && countVowels(input) >= 3 && hasConsecutiveLetters(input) {
		return true
	}
	return false
}
