package menu


import (
	"fmt"
	"os"

	"meca/cmd/internal/menu/menu_buttons"
)

var p = fmt.Print
var pln = fmt.Println
var fs = fmt.Fscan


func Fmenu() {
	pln("")
	pln("----------------")
	pln("Добро пожаловать в проект MECA.")
	pln("----------------")
	pln("Хотите прочесть zametku - введите 1")
	pln("Хотите записать свое собственное zametku - введите 2")
	pln("Хотите получить случайный zametku - введите 3")
	pln("Настройки MECA - 4")
	pln("Выйти - 0")
	p("Vvedite chislo: ")
	var user_choice int
	fs(os.Stdin, &user_choice)
	switch(user_choice) {
	   case 0:
			pln("Завершение работы программы")
			os.Exit(3)
	   case 1:
	     menu_buttons.Fmenu1()
			 Fmenu()
		 case 2:
			 menu_buttons.Fmenu2()
			 Fmenu()
		 case 3:
 			 menu_buttons.Fmenu3()
			 Fmenu()
		 case 4:
			 menu_buttons.Fmenu_settings()
			 Fmenu()
	   default:
	     pln("Ошибка. Попробуйте заново.")
	     Fmenu()
	}
}

