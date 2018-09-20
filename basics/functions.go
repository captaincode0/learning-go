package main

import "fmt"

func main(){
	/**
	 * A function that returns multiple values has to be used to make iterables
	 */
	
	w, x := make_values()

	fmt.Println(x)
	fmt.Println(w)

	fmt.Println(add(1,2,3,4,5,6,7,8,9,10)) 

	/**
	 * Closure
	 *
	 * Are functions inside of functions
	 */
	
	add2 := func(x, y int)int {
		return x+y
	}

	fmt.Println(add2(3,4))

	/**
	 * Defer
	 *
	 * Go has a special statement called defer which schedules a function call to be run afther function completes
	 */
	
	defer second()
	first()

	/**
	 * Panic & Recover
	 *
	 * Panic: cause a runtime error
	 * Recover: stops panic and returns the value
	 */
	//test_error()

	/**
	 * Defer 
	 *
	 * Panic immediately stops the execution of one function but can be paired with panic
	 */
	
	defer func() {
		str := recover()
		fmt.Println(str)
	}()

	panic("PANIC")
}	

func first(){
	fmt.Println("first")
}

func second(){
	fmt.Println("second")
}

// Need to describe one list of returns
func make_values() (int, int){
	return 5, 90 //and values are separated by comma
}

// Variadic functions
/**
 * There is a special form for the last parameter in Go function
 */
func add(args ...int)int {
	total := 0

	for _, v := range args {
		total += v
	}

	return total
}

func test_error(){
	panic("PANIC")
	str := recover()
	fmt.Println(str)
}