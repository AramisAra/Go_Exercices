package main

func print_base16() {
  var num rune = '0'
  var abf rune = 'a'

  for ; num <= '9'; num++ {
    _putchar(true, num, 1)
  }
  for ; abf <= 'f'; abf++ {
    _putchar(true, abf, 1)
  }
  _putchar(true, '\n', 1)
}
