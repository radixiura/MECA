package greeting_funcs

import (
	"fmt"
	"os"
)

func Hello() {
  fmt.Println("Hello, user!")
}

func Hello_usrans() {
  var user_answer int
  y := fmt.Fscan(os.Stdin, &user_answer)
  return y
}
