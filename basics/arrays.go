package main

import "fmt"

func main(){
	/**
	 * Array declaration
	 *
	 * var <array_name> [<size(int)>]int
	 */
	
	var x [5]int
	var total int = 0

	x[1] = 2

	fmt.Println(x[1])
	fmt.Println(x)

	/**
	 * A single _ (underscore) is used to tell compiler that does
	 */
	
	 for _, value := range x {
	 	total += value
	 }

	 fmt.Println(float32(total) / float32(len(x)))

	 /**
	  * Slices
	  *
	  * Slice is a segment of one array, and are indexables and have a length
	  *
	  * Arrays with missing length are created with length of zero
	  */
	 var y = []float64{32,55,56,44,55,89,32,78,46,65,}

	 /**
	  * To use one abreviated syntax, then use one [min:max] syntax,
	  * and create one delimited array
	  */

	  fmt.Println(y[2:6])

	 /**
	  * To slice one array just use make built-in function
	  */
	 fmt.Println(y[2])

	 z := make([]float64, 5, 10)

	 fmt.Println(z)

	 /**
	  * Go includes two built-in functions to assist with slices append and copy:
	  */
	 
	// Append function
	w := []int{1,2,3}
	w2 := make([]int, 2)

	fmt.Println(w)
	fmt.Println(w2)

	// Copy function
	copy(w2, w)
	fmt.Println(w2, w)

	/**
	 * Maps
	 *
	 * Are collections with key/value pairs
	 *
	 * var <map-name> map[<key datatype>]<value datatype>
	 *
	 * Maps need to be declared with make built-in function
	 */
	
	var my_map = make(map[string]int)

	my_map["0fffefa"] = 32
	fmt.Println(my_map["0fffefa"])

	my_map["a"] = 6

	/**
	 * Can also use built-in delete function
	 *
	 * When tries to get an inexistent index, then return zero value for given data type (int:0, string: "")
	 */
	delete(my_map, "0fffefa")

	fmt.Println(my_map);
}