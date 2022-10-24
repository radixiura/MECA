package log


import(
	"fmt"
	"meca/cmd/internal/log/log_funcs"
)


func Login() string {
	fmt.Println("LOGIN NOW")
  var user_login = log_funcs.Get_login()
  var user_password = log_funcs.Get_password()
  var user string
  user = user_login + user_password
  return user
}
