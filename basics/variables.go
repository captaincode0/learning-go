package main

import "fmt"

func main(){
	/**
	 * The syntax to declare one variable is
	 *
	 * var <variable_name> <datatype> = <value>
	 */
	
	var str string = "Hello world"

	fmt.Println(str)

	/**
	 * Multiple variable declaration
	 */
	
	var (
		a string = "data"
		b string = "data0"
		c string = "data1"
	)

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	/**
	 * Defining constants
	 *
	 * const <const_name> <datatype> = value
	 */
	
	const MX string = "MX"

	fmt.Println(MX)
}