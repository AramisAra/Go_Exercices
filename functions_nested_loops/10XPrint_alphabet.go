package main

func X10_printalpha() {
	abc := 'a'
	times := 0
	for ; times <= 9; times++ {
		for abc = 'a'; abc <= 'z'; abc++ {
			_putchar(abc)
		}
		_putchar('\n')
	}
}
