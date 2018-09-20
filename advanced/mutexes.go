package main

import (
	"fmt"
	"sync"
	"time"
)

func main(){
	/**
	 * A MUTEX (Mutual Exclusive Lock) locks a section of code to a single thread at a time and used to protect shared resources from non-atomic operations.l
	 */
	
	var m *sync.Mutex = new(sync.Mutex)

	for i := 0; i<10; i++{
		go func(i int){
			m.Lock()
			fmt.Println(i, "start")
			fmt.Println(time.Second)
			fmt.Println(i, "end")
			m.Unlock()
		}(i)
	} 

	var input string
	fmt.Scanln(&input)
}