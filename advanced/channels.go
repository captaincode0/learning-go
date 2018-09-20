package main

import (
	"fmt"
	"time"
)

func ping(c chan string){
	for i := 0; ; i++ {
		c <- "ping"
	}
}

func printer(c chan string){
	var msg string

	for { //use infinite loop
		msg = <- c
		fmt.Println(msg)
		time.Sleep(time.Second * 1)
	}
}

func main(){
	/**
	 * Channels provide a way for two goroutines to communicate another and sync their execution
	 *
	 * It's possible to pass a second parameter to make a function when creating a channel, indicates that channel is buffered with capacity of 1, both sides of channell will wait intul the other side is ready
	 *
	 * A BUFFERED CHANNEL IS ASYNC, RX/TX A MESSAGE WILL NOT WAIT UNLESS CHANNEL IS READY OR FULL
	 */
	var c chan string = make(chan string);

	go ping(c)
	go printer(c)

	var input string
	fmt.Scanln(&input)
	fmt.Println(input)

	/**
	 * Program will print "ping" forever, then a channel type is represented with the keyword chan followed by type of things that are passed on the channel, and '<-' operator is for pass channel selected type
	 */
}