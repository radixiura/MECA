package main


import (
	"fmt"
	"time"
	"os"
	"database/sql"

	"meca/cmd/internal/reg"
	"meca/cmd/internal/log"

	"meca/cmd/internal/menu"

	_ "github.com/lib/pq"
)


var p = fmt.Print
var pln = fmt.Println
var fs = fmt.Fscan


type product struct{
    id int
    model string
    company string
    price int
}




func main() {
		connStr := "user=postgres password=cordia dbname=productdb sslmode=disable"
		    db, err := sql.Open("postgres", connStr)
		    if err != nil {
		        panic(err)
		    } 
		    defer db.Close()
		     
		    result, err := db.Exec("insert into Products (model, company, price) values ('iPhone X', $1, $2)", 
		        "Apple", 72000)
		    if err != nil{
		        panic(err)
		    }
		    fmt.Println(result.LastInsertId())  // не поддерживается
		    fmt.Println(result.RowsAffected())  // количество добавленных строк
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
