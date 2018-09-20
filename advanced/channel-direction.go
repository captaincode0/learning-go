package main

import (
	"fmt"
	"time"
)

func pinger(c chan <- string){
	for i:=0; ; i++{
		c <- "ping"
	}
}

func ponger(c chan <- string){
	for i:=0; ; i++{
		c <- "pong"
	}
}

func printer(c <- chan string){
	var msg string

	for{
		msg = <- c

		fmt.Println(msg)
		time.Sleep(time.Second)
	}
}

func main(){
	/**
	 * A channel direction can be specified to receive/response data and what direction follows strictly
	 *
	 * <name> chan <- type :: To send data
	 * <name> <- chan type :: To send data
	 */
	var c chan string = make(chan string)

	go pinger(c)
	go ponger(c)
	go printer(c)

	var input string

	fmt.Scanln(&input)

	fmt.Println(input)
}