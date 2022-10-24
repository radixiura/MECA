package log_funcs


import (
 "fmt"
 "os"
)


var p = fmt.Print
var fs = fmt.Fscan

func Get_password() string {
  var user_password string
  p("Введите ваш пароль: ")
  fs(os.Stdin, &user_password)
  return user_password
}
