package main

func print_alphABET() {
  var ABC rune = 'a'
  
  for ; ABC <= 'z'; ABC++ {
    _putchar(true, ABC, 1)
  }

  ABC = 'A'

  for ; ABC <= 'Z'; ABC++ {
    _putchar(true, ABC, 1)
  }
  _putchar(true, '\n', 1)
}

