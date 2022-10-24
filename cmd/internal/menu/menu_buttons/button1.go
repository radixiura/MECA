package menu_buttons


import (
	"fmt"
	"os"

	"meca/cmd/internal/menu/menu_buttons/reading"
)


var p = fmt.Print
var pln = fmt.Println
var fs = fmt.Fscan


func Fmenu1() {
	pln("---------------+")
	pln("")
	pln("Выберите тему, в которой хотите прочесть zametku: ")
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
	        reading.Fmenu1_history()
		 case 2:
			reading.Fmenu1_psychology()
		 case 3:
 			reading.Fmenu1_philosophy()
		 case 4:
			reading.Fmenu1_math()
	     default:
	        break
	}
}
