package main

import "fmt"

func main() {
  for i := 0; i <= 100; i++ {
  if i % 3 == 0 {
   fmt.Println(i, "divisible by 3")
   } else {
    fmt.Println(i, "not divisible")
    }
   // Thi is good
  }
}
