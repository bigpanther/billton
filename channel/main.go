package main

import (
	"fmt"
	"time"
)

func main() {

	messages := make(chan string)
	var p = "ping"
	//	messages <- p
	go func() {
		messages <- p
	}()
	fmt.Println(<-messages)
	//buff()
	//work()
	pingpong()
}

func buff() {

	messages := make(chan string, 4)
	go func() {
		for i := 0; i < 1000; i++ {
			messages <- fmt.Sprintf("buffered%d", i)
			fmt.Println("entered value")
		}
	}()
	for {
		fmt.Println(<-messages)
		time.Sleep(time.Millisecond * 1000)
	}
}

func worker(done chan bool) {
	fmt.Print("working...")
	time.Sleep(time.Second)
	fmt.Println("done")

	fmt.Println(<-done)
}

func work() {

	done := make(chan bool, 1)
	go worker(done)

	done <- true
	time.Sleep(2 * time.Second)
}

func ping(pings chan<- string, msg string) {
	pings <- msg
}

func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}

func pingpong() {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	ping(pings, "passed message")
	pong(pings, pongs)
	fmt.Println(<-pongs)
}
