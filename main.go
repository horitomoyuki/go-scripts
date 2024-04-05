package main

import "fmt"

func main() {
  fmt.Printf("Hello Golang\nLet's Study Golang\n")
  sl := []int{1, 2, 3}
  if len(sl) > 2 {
    fmt.Println("unreachable code")
  }
}
