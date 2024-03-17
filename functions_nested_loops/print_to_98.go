package main

import "fmt"

func printTo98(n int) {
	if n > 98 {
		for ; n >= 98; n-- {
			fmt.Printf("%d, ", n)
		}
	} else if n < 98 {
		for ; n <= 98; n++ {
			fmt.Printf("%d, ", n)
		}
	} else {
		fmt.Printf("%d", n)
	}
	_putchar('\n')
}
