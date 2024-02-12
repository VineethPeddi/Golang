package main

import "fmt"

type CustomError struct {
	Msg string
}

func (ce *CustomError) error() string {
	return ce.Msg
}

func divideByZero(a, b int) *CustomError {
	if b == 0 {
		return &CustomError{
			Msg: "divide by zero error",
		}
	}
	return nil
}

func main() {
	// ce:=divideByZero(2,0)
	// fmt.Println(ce.error())
	defer func(){
		r := recover()
		fmt.Println("r: ", r)
	}()
	go func(){
		panic("panicking error")
	}()

}

func test() {
	panic("panicking error")
}

func st() {
	r := recover()
	fmt.Println("r: ", r)
}
