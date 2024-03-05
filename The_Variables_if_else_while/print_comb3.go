package main

func print_comb3() {
  num_1 := 0
  var num_2 int

  for ; num_1 < 10 ; num_1++ {
    for num_2 = num_1 + 1 ; num_2 < 10; num_2++{ 
      _putchar(false, '1', num_1)
      _putchar(false, '1', num_2) 
      if num_1 != 8 && num_2 != 10 {
         _putchar(true , ',', 1)
         _putchar(true , ' ', 1)
      }  
    }
  }
  _putchar(true, '\n', 1)
}
