package main

func print_comb() {
  num := '0'

  for ; num <= '9'; num++ {
    _putchar(true, num, 1)
    if num != '9' {
      _putchar(true, ',', 1)
      _putchar(true, ' ', 1)
    } else {
      _putchar(true, '$', 1)
    }
  }
  _putchar(true, '\n', 1)
}
