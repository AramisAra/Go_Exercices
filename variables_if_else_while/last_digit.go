package main

import (
  "fmt"
)
func last_digit() {
  R_number := RandInt(-30000, 30000)
  last := R_number % 10

  if last > 5 {
    fmt.Printf("Last digit of %d is %d and is greater than 5\n", R_number, last)
  } else if last < 6 && last != 0 {
    fmt.Printf("Last digit of %d is %d and is less then 6 and 0\n", R_number, last)
  } else {
    fmt.Printf("Last digit of %d is %d and is 0\n", R_number, last)
  }
}
