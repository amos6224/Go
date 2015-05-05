package main

import "fmt"

func main() {
  server := martini.Classic()
  server.Get("/", func() string {
    return "<h1>Hello, world</h1>"
  })

server.run

}
