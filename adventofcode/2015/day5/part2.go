package main

func isNice2(input string) bool {
	if !hasBadStrings(input, badStrings) && countVowels(input) >= 3 && hasConsecutiveLetters(input) {
		return true
	}
	return false
}
