package log_funcs


import (
 "os"
)


func Get_login() string {
  var user_login string
  p("Введите ваш login: ")
  fs(os.Stdin, &user_login)
  return user_login
}
