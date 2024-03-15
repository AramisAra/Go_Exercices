package main

import (
  "fmt"
  "math/rand"
)
func RandInt(lower, upper int) int {
	rng := upper - lower
	return rand.Intn(rng) + lower
}
func _putchar(arg1 bool, arg2 rune, arg3 int){
  if arg1{
    fmt.Printf("%c", arg2)
  } else {
    fmt.Printf("%d", arg3)
  }
}

func main() {
  fmt.Println("This is the first task")
  positive_or_negative()
  fmt.Println("This is the second task")
  last_digit()
  fmt.Println("This is the thrid task")
  print_alphabet()
  fmt.Println("This is the fourth task")
  print_alphABET()
  fmt.Println("This is the fifth task")
  print_alphabt()
  fmt.Println("This is the sixth task")
  print_numbers()
  fmt.Println("This is the seventh task")
  print_numberz()
  fmt.Println("This is the eighth task")
  print_tebahpla()
  fmt.Println("This is the ninth task")
  print_base16()
  fmt.Println("This is the tenth task")
  print_comb()
  fmt.Println("This is the first bonus task")
  print_comb3()
  fmt.Println("This is the second bonus task")
  print_comb4()
  fmt.Println("This is the third bonus task")
  print_comb5()
}
