package main

import "fmt"

func reset(x int){
	x = 0
}

func reset_with_pointer(x *int){
	*x = 0 //value on pointer need to be dereferenced
}

func main(){
	/**
	 * When a function is called and takes argument,
	 * this argument is copied into the function
	 */
	var x int = 5

	reset(x)

	fmt.Println(x)

	/**
	 * Then need to pass reference if want to change value of x
	 */
	reset_with_pointer(&x) //& takes the reference

	fmt.Println(x)

	/**
	 * Another way to get a pointer is to use built-in new function
	 */
	var y *int = new(int)
	*y = 4
	fmt.Println(*y)
	reset_with_pointer(y)
	fmt.Println(*y);
}