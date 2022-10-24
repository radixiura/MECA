package main


import (
	"fmt"
	"time"
	"os"

	"meca/cmd/internal/reg"
	"meca/cmd/internal/log"

	"meca/cmd/internal/menu"
)


var p = fmt.Print
var pln = fmt.Println
var fs = fmt.Fscan


func main() {
   Hello()
	 log.Login()
   menu.Fmenu()
}


func Hello() {
	pln("")
  pln("Привет! This is a MECA project.")
  pln("Вы уже имеете аккаунт?")
  p("Введите 0 если нет, 1 если да: ")
  var user_answer = User_answer()
  time.Sleep(2 * time.Second)
	switch(user_answer) {
        case 0:
						pln("")
            pln("Самое время зарегистрироваться!")
            reg.Registration()
        case 1:
						pln("")
        default:
						pln("")
            pln("Ошибка. Попробуйте заново.")
            Hello()
    }
}


func User_answer() int {
    var usr_ans int
    fs(os.Stdin, &usr_ans)
    return usr_ans
}
