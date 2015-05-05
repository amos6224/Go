// pointer is a type of v
package main
import "fmt"

type Salutation struct {
     name string
     greeting string}
func main() {
  var s = Salutation{"Jef", "Jen"}
  fmt.Println(s.name)
  fmt.Println(s.greeting)
}
