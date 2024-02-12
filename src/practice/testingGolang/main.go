package main

import "fmt"

func main() {
	testFn()
}
func testFn(){
	i:=0
	defer func(i int){fmt.Println(i)}(i)	
	i=110
	return
}

func handlingPanic(){
	defer func(){
		fmt.Println("in err1")
		if err:=recover();err!=nil{
			fmt.Println("err in recovery: ",err,"\n\n ")
		}
	}()
	defer func(){
		fmt.Println("in err2")
		if err:=recover();err!=nil{
			fmt.Println("err2 in recovery: ",err,"\n\n ")
		}
	}()
	type Student struct{
		Name string 
	}
	s:=&Student{}
	s=nil
	fmt.Println(s.Name)
}