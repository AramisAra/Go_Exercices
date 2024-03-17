package main

func timesTables() {
	var nums, divs, results int

	for nums = 0; nums < 9; nums++ {
		for divs = 0; divs < 10; divs++ {

			results = nums * divs
			if results > 9 {
				_putchar(results / 10)
				_putchar(results % 10)
			} else if divs != 0 {
				_putchar(' ')
				_putchar(results)
			} else {
				_putchar(results)
			}
			if divs != 9 {
				_putchar(',')
				_putchar(' ')
			}
		}
		_putchar('\n')
	}
}
