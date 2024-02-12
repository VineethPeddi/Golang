package main

import "fmt"

func main() {
	mp:=make(map[string]int)
	mp["a"]=1
	testMapUpdate(mp)
	fmt.Println(mp)
}

func testMapUpdate(mp map[string]int){
	mp["b"]=2
}

func test() func(string)int{

	return func(s string)int{return 1}
}