package main

func isalpha(c rune) int {
	if c >= 'a' && c <= 'z' {
		return 1
	} else if c >= 'A' && c <= 'Z' {
		return 1
	} else {
		return 0
	}
}
