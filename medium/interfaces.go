package main

import ("fmt";"math")

type Shape interface{
	area() float64
	perimeter() float64
}

type MultiShape struct{
	shapes []Shape
}

type Point struct{
	x float64
	y float64
}

func (p0 *Point) distance(p1 *Point)float64{
	return math.Sqrt((p1.x-p0.x)*(p1.x-p0.x)+(p1.y-p0.y)*(p1.y-p0.y))
}

type Circle struct{
	Point
	r float64
}

func (c *Circle) area()float64{
	return math.Pi*c.r*c.r
}

func (c *Circle) perimeter()float64{
	return 2*math.Pi*c.r
}

type Rectangle struct{
	origin 	Point
	end 	Point
}

func (r *Rectangle) area()float64{
	w := r.origin.distance(&Point{r.end.x, r.origin.y})
	l := r.origin.distance(&Point{r.origin.x, r.end.y})

	return w*l
}

func (r *Rectangle) perimeter()float64{
	w := r.origin.distance(&Point{r.end.x, r.origin.y})
	l := r.origin.distance(&Point{r.origin.x, r.end.y})

	return 2*w+2*l
}

func totalArea(m *MultiShape)float64{
	var area float64

	for _, s:=range m.shapes{
		area+=s.area()
	}

	return area
}

func main(){
	/**
	 * Interfaces are used to define super types for one or more types
	 * for example Shape could be a Circle, Rectangle, Triange, Hexagon, etc.
	 *
	 * And interface define a method set, are methods that need to define all sub-types
	 */
	var p Point = Point{0,0}
	var p1 Point = Point{3,4}

	fmt.Println(p)
	fmt.Println(p1)

	var distance float64 = p.distance(&p1)

	fmt.Println(distance)

	var rectangle *Rectangle = new(Rectangle)

	rectangle.origin = p
	rectangle.origin = p1

	var rarea float64 = rectangle.area()
	var rperimeter float64 = rectangle.perimeter()

	fmt.Println(rarea)
	fmt.Println(rperimeter)

	var circle *Circle = new(Circle)

	circle.r = 5

	fmt.Println(circle.area())
	fmt.Println(circle.perimeter())

	var shapes *MultiShape = new(MultiShape)

	shapes.shapes = []Shape{rectangle, circle}

	fmt.Println(totalArea(shapes))
}