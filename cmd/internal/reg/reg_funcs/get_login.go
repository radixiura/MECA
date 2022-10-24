package reg_funcs


import (
 "fmt"
 "os"
)


var p = fmt.Print
var pln = fmt.Println
var fs = fmt.Fscan


func Get_login() string {
  var new_user_login string
  p("Введите ваш новый логиH: ")
  fs(os.Stdin, &new_user_login)
  return new_user_login
}
