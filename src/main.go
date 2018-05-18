package main

import "fmt"

type Acter interface {
	act() string
}

type unitState struct {
	a Acter
}

type forReal struct {}

func(r forReal) act()  string{
	return "real"
}
func act(string2 string)  {

}

//если мы не хотим менять что-то в структуре, то без указателя
//	  this calls receiver
//			|
func (u unitState) f(p string) string  {
	s:= u.a.act()
	return p + ":" + s
}

//и её можно в инит присваивать какой-то, чтобы даже в хендлере не менять сигнатуру
//var t func() string
//
//func mock() string {
//	return "0"
//}
//
//func real()  string{
//	return "real"
//}
//
//func f(p string) string  {
//	s:=t()
//	return p + ":" + s
//}
//
//func main(){
//
//	t=mock//or t=real
//	fmt.Print(f("text"))
//
//	// now lets try using interface
//
//
//
//}

func main() {
	var ob = unitState{forReal{}}
	fmt.Print(ob.f("text"))
}
