package main

import (
	"time"
	"fmt"
	"os"
)

func main() {
   Fhello()
   Flogin()
   Fmenu()
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
            fmt.Println("Ошибка. Попробуйте заново.")
            Fhello()
    }
}


func Fget_user_answer() int {
    var usr_ans int
    fmt.Fscan(os.Stdin, &usr_ans)
    return usr_ans
}


func Fregistration() {
  var  fuser_login = Fregistration_get_login()
  fmt.Print(fuser_login) //raspechatyvaet pervoe vvedennoe!
  var fuser_password = Fregistration_get_password()
  fmt.Print(fuser_password)
}



func Flogin() {
  var luser_login = Flogin_get_login()
  fmt.Print(luser_login) //raspechatyvaet pervoe vvedennoe!
  var luser_password = Flogin_get_password()
  fmt.Print(luser_password)
}


func Fregistration_get_login() string {
  var new_user_login string
  fmt.Print("Введите ваш новый логин. On doljen soderjat bolee 8mi simvolov: ")
  fmt.Fscan(os.Stdin, &new_user_login)
  l := len(new_user_login)
  fmt.Println(l)
  if l <= 8 {
    time.Sleep(3 * time.Second)
  	fmt.Println("login doljen soderjat bolee 8mi simvolov. Poprobuyte eshe raz")
  	Fregistration_get_login()
  } else if l >= 8 {
  	fmt.Println("zbs")
  } else {
  	fmt.Print("Ошибка. Попробуйте заново.")
  }
  return new_user_login
}



func Fregistration_get_password() string {
  var new_user_password string
  fmt.Print("Введите ваш новый parol. On doljen soderjat bolee 8mi simvolov: ")
  fmt.Fscan(os.Stdin, &new_user_password)
  l := len(new_user_password)
  fmt.Println(l)
  if l <= 8 {
    time.Sleep(3 * time.Second)
  	fmt.Println("Parol doljen soderjat bolee 8mi simvolov. Poprobuyte eshe raz")
  	Fregistration_get_password()
  } else if l >= 8 {
  	fmt.Println("zbs")
  } else {
  	fmt.Print("Ошибка. Попробуйте заново.")
  }
  return new_user_password
}


func Flogin_get_login() string {
  var user_login string
  fmt.Print("Vvedite vash login")
  fmt.Fscan(os.Stdin, &user_login)
  fmt.Println(user_login)
  return user_login
}



func Flogin_get_password() string {
  var user_password string
  fmt.Print("Vvedite vash login")
  fmt.Fscan(os.Stdin, &user_password)
  fmt.Println(user_password)
  return user_password
}




func Fmenu() {
	fmt.Println("Добро пожаловать в проект MECA.")
	fmt.Println("Хотите прочесть воспоминание - введите 1")
	fmt.Println("Хотите записать свое собственное воспоминание - введите 2")
	fmt.Println("Хотите получить случайный осколок памяти - введите 3")
	fmt.Println("Настройки MECA - 4")
	fmt.Println("Выйти - 0")
	var user_choice int
	fmt.Fscan(os.Stdin, &user_choice)
	switch(user_choice) {
	     case 0:
				 Fmenu_escape()
				 os.Exit(3)
	     case 1:
	        Fmenu1()
			 case 2:
					Fmenu2()
			 case 3:
 					Fmenu3()
			 case 4:
					Fmenu_settings()
	     default:
	        fmt.Println("Ошибка. Попробуйте заново.")
	        Fhello()
				}
}

func Fmenu_escape() {
	fmt.Println("Завершение работы программы")
	os.Exit(3)
}


func Fmenu1() {
	fmt.Println("Choose the theme you want to read about: ")
	fmt.Println("You want to read - istoria")
	fmt.Println("You want to read - psiholog")
	fmt.Println("You want to read filosof")
	fmt.Println("You want to read - fizmat")
	fmt.Println("You want to nazad")
	var user_choice int
	switch(user_choice) {
	     case 0:
	        Fmenu()
	     case 1:
	        Fmenu1_history()
			 case 2:
					Fmenu1_psychology()
			 case 3:
 					Fmenu1_philosophy()
			 case 4:
					Fmenu1_fizmat()
	     default:
	        fmt.Println("Oshibka. Poprobuyte zanovo!")
	        Fmenu()
				}
}


func Fmenu1_history() {
	fmt.Println("")
	}


func Fmenu1_psychology() {
	fmt.Println("")
}


func Fmenu1_philosophy() {
	fmt.Println("")
}


func Fmenu1_fizmat() {
	fmt.Println("")
}


func Fmenu2() {
	fmt.Println("Choose the theme you want to write about: ")
	fmt.Println("You want to write - istoria")
	fmt.Println("You want to write - psiholog")
	fmt.Println("You want to write filosof")
	fmt.Println("You want to write - fizmat")
	fmt.Println("You want to nazad")
	var user_choice int
	switch(user_choice) {
	     case 0:
	        Fmenu()
	     case 1:
	        Fmenu2_history()
			 case 2:
					Fmenu2_psychology()
			 case 3:
 					Fmenu2_philosophy()
			 case 4:
					Fmenu2_fizmat()
	     default:
	        fmt.Println("Oshibka. Poprobuyte zanovo!")
	        Fmenu()
				}
}


func Fmenu2_history() {
	fmt.Println("")
	}


func Fmenu2_psychology() {
	fmt.Println("")
}


func Fmenu2_philosophy() {
	fmt.Println("")
}


func Fmenu2_fizmat() {
	fmt.Println("")
}



func Fmenu3() {
	fmt.Println("random note")
}

func Fmenu_settings() {
	fmt.Println("Settings")
}
