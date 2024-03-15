package main

import (
  "fmt"
)

func print_numbers() {
  var num int = 0

  for ; num <= 9; num++ {
    fmt.Printf("%d", num)
  }
  fmt.Printf("\n")
}
