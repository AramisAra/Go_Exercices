package main

func print_comb4() {
  var numa int
  var numb int
  var numc int

  for numa = 0; numa < 10; numa++ {
    for numb = numa + 1; numb < 10; numb++ {
      for numc = numb + 1; numc < 10; numc++ {
        _putchar(false, '1', numa)
        _putchar(false, '1', numb)
        _putchar(false, '1', numc)
        if numa != 7 && numb != 9 && numc != 10 {
          _putchar(true, ',', 1)
          _putchar(true, ' ', 1)
        }
      }
    }
  }
  _putchar(true, '\n', 1)
}
