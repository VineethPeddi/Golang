package main

import (
	"fmt"
	"time"
)

func main() {
	// basicGoroutineExample()
	// creatingWorkers()
	// ch:=make(chan int,2)
	// fmt.Println("about to send")
	// go test(ch)
	// time.Sleep(3*time.Second)
	// ch<-3
	// time.Sleep(3*time.Second)
	// ch<-4
	// ch<-5
	// fmt.Println("after send")

	fmt.Println("In main")

	go func ()  {
		fmt.Println("In gr1")
		go func(){
			fmt.Println("In inner goroutine before sleep")
			time.Sleep(3*time.Second)
			fmt.Println("In inner goroutine after sleep")
		}()
		fmt.Println("exiting from gr1")
	}()

	fmt.Println("about to sleep in main")
	time.Sleep(5*time.Second)
	fmt.Println("about to exit in main")
}

func test(ch<- chan int){
	fmt.Println("waiting for data from channel")
	val:=<-ch
	fmt.Println("val: ",val)
}

func creatingWorkers() {
	in := make(chan int)
	out := make(chan int)
	// go readMultiplyAndSend(in, out, "four")

	go func() {
		in <- 1
		in <- 2
		in <- 3
		fmt.Println("about to input 4")
		// time.Sleep(3*time.Second)
		in <- 4
	}()

	
	go readMultiplyAndSend(in, out, "first")
	go readMultiplyAndSend(in, out, "second")
	go readMultiplyAndSend(in, out, "third")
	
	fmt.Println(<-out)
	fmt.Println(<-out)
	fmt.Println(<-out)
	fmt.Println("Waiting for 4 output")
	fmt.Println(<-out)
}

func readMultiplyAndSend(in <-chan int, out chan<- int, channelNumber string) {
	for {
		val := <-in
		fmt.Println("val: ", val, " in ", channelNumber, " channel")
		res := val * 2
		out <- res
		fmt.Println("after sending data from ",channelNumber," channel")
	}
}






func basicGoroutineExample() {
	c := make(chan int)
	go multiply(4, c)
	fmt.Println("waiting for data")
	fmt.Println(<-c)
	fmt.Println("received data")
}

func multiply(num int, c chan int) {
	time.Sleep(3 * time.Second)
	c <- num * 2
	time.Sleep(3 * time.Second)
	fmt.Println("sent data to channel")
	fmt.Println("about to return")
}
