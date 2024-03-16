package main

func printLast(n int) int {
	n = n % 10
	n = _abs(n)
	_putchar(n)
	return n
}
