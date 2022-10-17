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
	fmt.Println("")
  fmt.Println("Привет! This is a MECA project.")
  fmt.Println("Вы уже имеете аккаунт?")
  fmt.Print("Введите 0 если нет, 1 если да: ")
  var user_answer = Fget_user_answer()
	switch(user_answer) {
        case 0:
						fmt.Println("")
            fmt.Println("Самое время зарегистрироваться!")
            Fregistration()
        case 1:
						fmt.Println("")
        default:
						fmt.Println("")
            fmt.Println("Ошибка. Попробуйте заново.")
            Fhello()
    }
}


func Fget_user_answer() int {
    var usr_ans int
    fmt.Fscan(os.Stdin, &usr_ans)
    return usr_ans
}

//----//

func Fregistration() {
    var ruser_inf string = Fregistration_get_name()
		fmt.Println("")
		fmt.Println("registered user login+password is:", ruser_inf)
		time.Sleep(3 * time.Second)
}

func Flogin() {
    var luser_inf string = Flogin_get_name()
		fmt.Println("")
		fmt.Println("logged in user login+password is:", luser_inf)
		time.Sleep(3 * time.Second)
}

//----//

func Fregistration_get_name() string {
  var ruser_login = Fregistration_get_name_login()
  var ruser_password = Fregistration_get_name_password()
	var ruser string
	ruser = ruser_login + ruser_password
	return ruser
}

func Flogin_get_name() string {
  var luser_login = Flogin_get_name_login()
  var luser_password = Flogin_get_name_password()
	var luser string
	luser = luser_login + luser_password
	return luser
}

//----//

func Fregistration_get_name_login() string {
  var new_user_login string
  fmt.Print("Введите ваш новый логин. Он должен содержать более 8ми символов: ")
  fmt.Fscan(os.Stdin, &new_user_login)
	// Proverka dliny
  l := len(new_user_login)
  if l <= 8 {
    time.Sleep(3 * time.Second)
  	fmt.Println("Логин должен содержать более 8ми символов. Попробуйте еще раз")
  	Fregistration_get_name_login()
  } else if l >= 8 {
  	fmt.Println("")
  } else {
  	fmt.Print("Ошибка. Попробуйте заново.")
  }

  return new_user_login
}

func Fregistration_get_name_password() string {
  var new_user_password string
  fmt.Print("Введите ваш новый пароль. Он должен содержать более 8ми символов: ")
  fmt.Fscan(os.Stdin, &new_user_password)
	// Proverka dliny
  l := len(new_user_password)
  if l <= 8 {
    time.Sleep(3 * time.Second)
  	fmt.Println("Пароль должен содержать более 8ми символов. Попробуйте еще раз: ")
  	Fregistration_get_name_password()
  } else if l >= 8 {
  	fmt.Println("")
  } else {
  	fmt.Print("Ошибка. Попробуйте заново.")
  }

  return new_user_password
}

//----//

func Flogin_get_name_login() string {
  var user_login string
  fmt.Print("Введите ваш логин: ")
  fmt.Fscan(os.Stdin, &user_login)
  return user_login
}

func Flogin_get_name_password() string {
  var user_password string
  fmt.Print("Введите ваш пароль: ")
  fmt.Fscan(os.Stdin, &user_password)
  return user_password
}

//----//

func Fmenu() {
	fmt.Println("")
	fmt.Println("----------------")
	fmt.Println("Добро пожаловать в проект MECA.")
	fmt.Println("----------------")
	fmt.Println("Хотите прочесть zametku - введите 1")
	fmt.Println("Хотите записать свое собственное zametku - введите 2")
	fmt.Println("Хотите получить случайный zametku - введите 3")
	fmt.Println("Настройки MECA - 4")
	fmt.Println("Выйти - 0")
	fmt.Println("")
	fmt.Print("Vvedite chislo: ")
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

//----//

func Fmenu_escape() {
	fmt.Println("Завершение работы программы")
	os.Exit(3)
}


func Fmenu1() {
	fmt.Println("")
	fmt.Println("Выберите тему, в которой хотите прочесть zametku: ")
	fmt.Println("История - 1")
	fmt.Println("Психология - 2")
	fmt.Println("Философия - 3")
	fmt.Println("Физмат - 4")
	fmt.Println("Назад - 0")
	fmt.Println("")
	fmt.Print("Vvedite chislo: ")
	var user_choice int
	fmt.Fscan(os.Stdin, &user_choice)
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
	        fmt.Println("Ошибка. Попробуйте заново.")
	        Fmenu()
				}
}

//----//

func Fmenu1_history() {
	fmt.Println("")
	fmt.Println("Британия и США начали войну из‑за убийства свиньи.")
	Fmenu()
	}

func Fmenu1_psychology() {
	fmt.Println("")
	fmt.Println("Нет для человека лучшего звука, чем его произнесенное имя.")
	Fmenu()
}

func Fmenu1_philosophy() {
	fmt.Println("")
	fmt.Println("Интроспекция (от лат. introspecto — смотрю внутрь) — способ самопознания, в ходе которого человек наблюдает за своей внутренней реакцией на события внешнего мира.")
	Fmenu()
}

func Fmenu1_fizmat() {
	fmt.Println("")
	fmt.Println("В переводе с арабского «цифра» слово означает «ноль», но исторически так сложилось, что этим словом мы называем вообще все цифры.")
	Fmenu()
}

//----//

func Fmenu2() {
	fmt.Println("")
	fmt.Println("Выберите тему, в которой хотите записать zametku")
	fmt.Println("История - 1")
	fmt.Println("Психология - 2")
	fmt.Println("Философия - 3")
	fmt.Println("Физмат - 4")
	fmt.Println("Назад - 0")
	fmt.Println("")
	fmt.Print("Vvedite chislo: ")
	var user_choice int
	fmt.Fscan(os.Stdin, &user_choice)
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
	        fmt.Println("Ошибка. Попробуйте заново.")
	        Fmenu()
				}
}

//----//

func Fmenu2_history() {
	fmt.Println("")
	fmt.Println("Vvedite zametku")
	Fmenu()
	}

func Fmenu2_psychology() {
	fmt.Println("")
	Fmenu()
}

func Fmenu2_philosophy() {
	fmt.Println("")
	Fmenu()
}

func Fmenu2_fizmat() {
	fmt.Println("")
	Fmenu()
}

//----//

func Fmenu3() {
	fmt.Println("")
	fmt.Println("Случайное zametka: ")
	fmt.Println("")
	fmt.Println("Петрикор - запах земли после дождя")
	fmt.Println("")
	Fmenu()
}

//----//

func Fmenu_settings() {
	fmt.Println("")
	fmt.Println("Настройки")
	Fmenu()
}
