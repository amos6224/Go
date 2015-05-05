package main

import "fmt"

func main() {
fmt.Print("Enter Celsius Number: ")
var input float64
fmt.Scanf("%f", &input)

output := input * 5/9

fmt.Println(output)
}
