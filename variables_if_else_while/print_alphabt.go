package main

func print_alphabt() {
  var abc rune = 'a'

  for ; abc <= 'z'; abc++ {
    if abc != 'e' && abc != 'q'{
      _putchar(true, abc, 1)
    }
  }
  _putchar(true, '\n', 1)
}
