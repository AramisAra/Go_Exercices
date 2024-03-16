package main

func _abs(n int) int {
	if n < 0 {
		n = -1 * n
		return n
	} else {
		return n
	}
}
