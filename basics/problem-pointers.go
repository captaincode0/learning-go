package main

import "fmt"

func swap(x *int, y *int){
	var z int = *x
	*x = *y
	*y = z
}

func main(){
	/**
	 * Write a program that can swap two integers (x:1, y:2) swap(&x, &y), should give (x:2,y:1)
	 */
	var (
		x int = 1
		y int = 2
	)

	swap(&x, &y)

	fmt.Println(x, ":", y)
}