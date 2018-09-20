package main

import ("fmt";"math")

type Circle struct{
	x,y,r float64
}

func circleArea(c Circle)float64{
	return math.Pi*c.r*c.r
}

func main(){
	/**
	 * A struct can be initialized as following
	 *
	 * var c Circle
	 *
	 * Using pointers
	 *
	 * var c *Circle = new(Circle)
	 *
	 * c.x = 0
	 * c.y = 0
	 * c.r = 10
	 *
	 * Assigment by creation
	 *
	 * var c Circle = Circle{x: 0, y:0, r:5}
	 * var c Circle = Circle{0, 0, 5}
	 */
	
	var c Circle = Circle{0,0,5}

	fmt.Println(circleArea(c))
}