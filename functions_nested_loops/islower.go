package main

func islower(c rune) int {
	if c >= 'a' && c <= 'z' {
		return 1
	} else {
		return 0
	}
}
