package main

func jack_banuer() {
	Hour := 0
	mins := 0

	for ; Hour < 24; Hour++ {
		for mins = 0; mins < 60; mins++ {
			if Hour < 10 {
				_putchar(0)
				_putchar(Hour)
			} else {
				_putchar(Hour)
			}
			_putchar(':')
			if mins < 10 {
				_putchar(0)
				_putchar(mins)
			} else {
				_putchar(mins)
			}
			_putchar('\n')
		}
	}
}
