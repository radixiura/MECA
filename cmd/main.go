package main

import (
	"fmt"
	"os"
	"time"

	"meca/cmd/internal/log"
	"meca/cmd/internal/reg"

	"meca/cmd/internal/menu"

	_ "github.com/lib/pq"
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
	var userAnswer = userAnswer()
	time.Sleep(2 * time.Second)
	switch userAnswer {
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

func userAnswer() int {
	var usrAns int
	fs(os.Stdin, &usrAns)
	return usrAns
}
