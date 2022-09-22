package main

import (
	"fmt"
	"time"
)

func worker(d chan int) {
	fmt.Println("Working....Done")
	time.Sleep(time.Second * 2)
	d <- 6
}
func main() {

	fmt.Println("This is Main Function")
	mess := make(chan int, 1)

	// mess <- 2
	// go func() { mess <- 4 }()
	go worker(mess)
	fmt.Println("This is the msg val : ", <-mess)
	pings := make(chan int, 1)
	pongs := make(chan int, 1)

	ping(pings, 21)
	pong(pings, pongs)

	fmt.Println("This is the Pong value : ", <-pongs)

	timer1 := time.NewTimer(2 * time.Second)

	fmt.Println("Timer1 Val", <-timer1.C)

}

func ping(ping chan<- int, s int) {
	ping <- s

}

func pong(ping <-chan int, pong chan<- int) {
	pong <- <-ping
}
