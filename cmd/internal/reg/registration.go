package reg


import(
	"meca/cmd/internal/reg/reg_funcs"
)


func Registration() string {
  var new_user_login = reg_funcs.Get_login()
  var new_user_password = reg_funcs.Get_password()
  var new_user string
  new_user = new_user_login + new_user_password
  return new_user
}
