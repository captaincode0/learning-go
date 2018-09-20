package main

import ("fmt";"math")

type Circle struct{
	x,y,r float64
}

/**
 * This is one method for one struct, that can be used and handled to wtrit
 */
func (c *Circle) area()float64{
	return math.Pi*c.r*c.r
}

func main(){
	var c *Circle = new(Circle)

	c.r = 5

	fmt.Println(c.area())
}