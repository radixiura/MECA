package main
import (
    "fmt"
    "./funcs"
)
 
func main() {
  fmt.Println("Привет!")
  var x int = funcs.Greeting()
  fmt.Println(x)
}
