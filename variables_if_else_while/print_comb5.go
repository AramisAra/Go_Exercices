package main

func print_comb5() {
	var num_a int
	var num_b int

	for num_a = 0; num_a < 99; num_a++ {
		for num_b = num_a + 1; num_b < 100; num_b++ {
			_putchar(false, '1', (num_a / 10))
			_putchar(false, '1', (num_a % 10))
			_putchar(true, ' ', 1)
			_putchar(false, '1', (num_b / 10))
			_putchar(false, '1', (num_b % 10))

			if num_a != 98 || num_b != 99 {
				_putchar(true, ',', 1)
				_putchar(true, ' ', 1)
			}
		}
	}
	_putchar(true, '\n', 1)
}
