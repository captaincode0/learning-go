package main

import "fmt"

func main(){
	/**
	 * For Loop
	 *
	 * for <variables>; <condition>; increment {
	 * 
	 * }
	 */
	
	 for i := 1; i<=10; i++ {
	 	if i%2 == 0 {
	 		fmt.Println(i, " even")
	 	} else {
	 		fmt.Println(i, " odd")
	 	}
	 }
}