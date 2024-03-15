package main

func print_comb5() {
  var numa int
  var numb int

  for numa = 0; numa < 99; numa++ {
    for numb = numa + 1; numb < 100; numb++ {
      _putchar(false, '1', (numa / 10))
      _putchar(false, '1', (numa % 10))
      _putchar(true, ' ', 1)
      _putchar(false, '1', (numb / 10))
      _putchar(false, '1', (numb % 10))

      if (numa != 98 || numb != 99) {
        _putchar(true, ',', 1)
        _putchar(true, ' ', 1)
      }
    }
  }
  _putchar(true, '\n', 1)
}
