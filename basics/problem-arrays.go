package main

import "fmt"

func main(){
	/**
	 * Writes a small program that finds the smaller number in this list
	 */
	
	var x = []int{
		48, 96, 86, 68,
		57, 82, 63, 70,
		37, 34, 83, 27,
		19, 97,  9, 17,
	}

	var array_size int = len(x)

	var small_number int = x[0]

	for i:=1; i<array_size; i++ {
		if small_number > x[i] {
			small_number = x[i]
			continue
		}
	}

	fmt.Println(small_number)

	/**
	 * When use a for as following way
	 *
	 * for _, v := range xs {
	 * 		//iterates will xs is iterated at all
	 * }
	 */
}
