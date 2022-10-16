package main

import (
	"fmt"
	"os"
)

func main() {
   Fhello()
   Flogin()
}


func Fhello() {
  fmt.Println("Привет!")
  fmt.Print("Вы уже имеете аккаунт?: ")
	fmt.Print("Введите 0 если нет, 1 если да: ")
  var user_answer = Fget_user_answer()
	switch(user_answer) {
        case 0:
            fmt.Println("Самое время зарегистрироваться!")
            Fregistration()
        case 1:
            fmt.Println("Отлично.")
        default:
            fmt.Println("Oshibka. Poprobuyte zanovo!")
            Fhello()
    }
}


func Fget_user_answer() int {
    var usr_ans int
    fmt.Fscan(os.Stdin, &usr_ans)
    return usr_ans
}


func Fregistration() {
  var new_user_login, new_user_password string
  fmt.Print("Введите ваш новый логин: ")
  fmt.Fscan(os.Stdin, &new_user_login)
  fmt.Print("Введите ваш новый пароль: ")
  fmt.Fscan(os.Stdin, &new_user_password)
}


func Flogin() {
  var user_login, user_password string
  fmt.Print("Введите ваш логин: ")
  fmt.Fscan(os.Stdin, user_login)
  fmt.Print("Введите ваш пароль: ")
  fmt.Fscan(os.Stdin, user_password)
}
