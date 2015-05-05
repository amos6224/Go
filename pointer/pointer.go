package main

import "fmt"

func main() {
  message := "hello World!"

  var greeting *string = &message
  // manipulate the message value
  *greeting = "hi"


  fmt.Println(message, greeting, *greeting)
}
