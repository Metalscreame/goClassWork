package main

import (
	"fmt"
	"time"
)

func add(a, b int) int {
	return a + b
}

func logWrapper(f func(a, b int) int) func(a, b int) int {

	return func(a, b int) int {
		fmt.Println("---------DECORATED BEFORE-----------")
		d := f(a, b)
		fmt.Println("---------DECORATED AFTER-----------")
		return d
	}

}

func logTimedWrapper(f func(a, b int) int) func(a, b int) int {
	return func(a, b int) int {
		fmt.Println("---------Time decorator-----------")
		start := time.Now()

		d := f(a, b)

		end:= time.Now()
		fmt.Printf("It took %v nanoseconds",start.Sub(end).Nanoseconds())

		fmt.Println()
		return d
	}
}

func main() {
	fmt.Printf("%d\n", add(1, 2))

	loggerAdd := logWrapper(add)
	loggerTimedAdd := logTimedWrapper(add)

	fmt.Println(logWrapper(add)(1, 2)) // logWrapper returns func, thats why we make this syntax
	fmt.Println(loggerAdd(1, 2))

	fmt.Println(loggerTimedAdd(1, 2))


	loggedFull :=logTimedWrapper(loggerAdd) // можно делать раппер на рапер
	fmt.Println(loggedFull(1, 2))

fmt.Println()
	//Замыкания. высший пилотаж функционального прогерства
	var x int
	addClosure:= accumulatorCreator(0)


	x=addClosure(1)
	fmt.Println(x)
	x=addClosure(10)
	fmt.Println(x)
	x=addClosure(-5)
	fmt.Println(x)


	fmt.Println()
	fib:=fibCreator()

	fmt.Println(fib())
	fmt.Println(fib())
	fmt.Println(fib())
	fmt.Println(fib())
	fmt.Println(fib())
	fmt.Println(fib())
	fmt.Println(fib())


}

func accumulatorCreator(initVal int) func(int)int {
	sum:=initVal
	return func(a int) int {
		sum+= a
		return sum
	}

}

func fibCreator()  func()int{
	n1,n2:=1,1
	return func( ) int {
		res:=n2
		n1,n2=n2,n2+n1
		return res
	}

}
