package main

import "fmt"

func main() {
  for i := 0; i <= 100; i ++{
  if i % 3 == 0{
   fmt.Println("Fizz")
  }
   if i % 5 == 0{
   fmt.Println("Buzz")
   }
    }
}
