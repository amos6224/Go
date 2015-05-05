package main
import "fmt"
//declare a new type
type Salutation string
func main() {
    //var  message string = "Hello Jeff"
    //var greeting *string = &message
    //*greeting = "hi"
    var message Salutation = "Hello Jeff"
    fmt.Println(message)
    //fmt.Println(message, greeting)
  }
