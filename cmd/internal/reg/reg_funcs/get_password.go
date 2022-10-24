package reg_funcs


import (
 "os"
 "time"
)


func Get_password() string {
  var new_user_password string
  p("Введите ваш новый пароль. Он должен содержать более 8ми символов: ")
  fs(os.Stdin, &new_user_password)
	// Proverka dliny
  l := len(new_user_password)
  if l <= 8 {
    time.Sleep(3 * time.Second)
  	pln("Пароль должен содержать более 8ми символов. Попробуйте еще раз: ")
  	Get_password()
  } else if l >= 8 {
  	pln("")
    time.Sleep(3 * time.Second)
  }
  return new_user_password
}
