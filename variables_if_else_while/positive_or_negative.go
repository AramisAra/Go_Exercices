package main

import (
  "fmt"
)

func positive_or_negative() {
  R_number := RandInt(-30000, 30000)
  if R_number > 0 {
    fmt.Printf("%d is positive\n", R_number)
  } else if R_number < 0 {
    fmt.Printf("%d is negative\n", R_number)
  } else {
    fmt.Printf("%d is zero\n", R_number)
  }
}
