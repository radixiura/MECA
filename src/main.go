package main

import (
	"fmt"
	"os"
)

func main() {

    var am = Fhero()  
    fmt.Println(am)
  
   Fhello()
}


func Fhello() {
  fmt.Println("Hello, ")
	fmt.Print("Vy uzhe imeete acc?: ")
	fmt.Print("Vvedite 0 esli net, 1 esli da: ")
  var user_answer int
	fmt.Fscan(os.Stdin, &user_answer)
	switch(user_answer) {
        case 0:
            fmt.Println("Samoe vremya zaregatsa")
            Fregistration()
        case 1:
            fmt.Println("Otlichno. Vvedite log i parol")
            Flogin()
        default:
            fmt.Println("Oshibka. Poprobuyte zanovo!")
            Fhello()
    }
}


func Fhero(z int) int {
    var z int
    fmt.Print("Vvedite vash noviy login: ")
    fmt.Fscan(os.Stdin, &z)
    return z
}


func Fregistration() {
  var new_user_login string
  fmt.Print("Vvedite vash noviy login: ")
  fmt.Fscan(os.Stdin, &new_user_login)
}

func Flogin() {
  var user_login string
  fmt.Print("Vvedite vash login: ")
  fmt.Fscan(os.Stdin, user_login)
}
