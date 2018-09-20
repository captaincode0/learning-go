package main

import (
	"fmt"
	"time"
	"math/rand"
)

func f(n int){
	for i:=1;i<=n;i++{
		fmt.Println(n,":",i)
		amt := time.Duration(rand.Intn(250))
		time.Sleep(time.Millisecond * amt)
	}
}

func main(){
	/**
	 * Large programs are made of smaller sub-programs, for 
	 * example a web server that handles requests for web-browsers and serves HTML webpages in response, each request is handled as a small program.
	 */
	
	/**
	 * Go Routines
	 *
	 * A goroutine is function capable of running concurrently with other functions, and to create a go routine use the keyword go followed by a function invocation
	 */

	go f(10)

	var input string
	fmt.Scanln(&input)

	fmt.Println(input)
}