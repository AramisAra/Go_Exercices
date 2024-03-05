package main

import (
	"fmt"
)

func _putchar(arg1 bool, arg2 rune, arg3 int){
  if arg1{
    fmt.Printf("%c", arg2)
  } else {
    fmt.Printf("%d", arg3)
  }
}

func main() {
	fmt.Println("This is the first task")
	putchar()
}