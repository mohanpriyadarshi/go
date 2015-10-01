// interface project main.go
package main

import (
	"fmt"
)
type Men interface {
    SayHi()
    Sing(lyrics string)
}

type Human struct {
    name string
    age int
    phone string
}
type Student struct {
  //  Human //an anonymous field of type Human
    school string
    loan float32
}


func (h Human) SayHi() {
    fmt.Printf("Hi, I am %s you can call me on %s\n", h.name, h.phone)
}

//A human can sing a song
func (h Human) Sing(lyrics string) {
    fmt.Println("La la la la...", lyrics)
}

func (h Student)SayHi(){
	fmt.Println("student data structer",h.school)
}
func (h Student)Sing(l string){
	fmt.Println("student data structer",l)
}



func main() {
	//mike := Human{"Mike", 25, "222-222-XXX"}
   var i Men
    //i can store a Student
    i = Human{"Mike", 25, "222-222-XXX"}
    fmt.Println("This is Mike, a Student:")
    i.SayHi()
   i.Sing("November rain")
	i =Student{"Vimanapuara",0.00}
	i.SayHi()
	i.Sing("ass")
}
