package main

import ("fmt";"time")

func main(){
	var c1 chan string = make(chan string)
	var c2 chan string = make(chan string)

	go func(){
		for{
			c1 <- "from 1"
			time.Sleep(time.Second)
		}
	}()

	go func(){
		for{
			c2 <- "from 2"
			time.Sleep(time.Second)
		}
	}()

	go func(){
		for{
			select {
				case msg1 := <- c1:
					fmt.Println(msg1)
				case msg2 := <- c2:
					fmt.Println(msg2)
				case <- time.After(time.Second):
					fmt.Println("timeout")
				default:
					fmt.Println("nothing ready")
			}
		}
	}()

	var input string
	fmt.Scanln(&input)
	fmt.Println(input)
}