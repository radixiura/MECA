package menu_buttons


import (
	"os"
	"meca/cmd/internal/menu/menu_buttons/writing"
)


func Fmenu2() {
	pln("---------------+")
	pln("")
	pln("Выберите тему, в которой хотите zapisat zametku: ")
	pln("История - 1")
	pln("Психология - 2")
	pln("Философия - 3")
	pln("Физмат - 4")
	pln("Назад - 0")
	pln("")
	p("Vvedite chislo: ")
	var user_choice int
	fs(os.Stdin, &user_choice)
	switch(user_choice) {
	     case 0:
	        break
	     case 1:
	        writing.Fmenu2_history()
		 case 2:
			writing.Fmenu2_psychology()
		 case 3:
 			writing.Fmenu2_philosophy()
		 case 4:
			writing.Fmenu2_math()
	     default:
	        break
	}
}
