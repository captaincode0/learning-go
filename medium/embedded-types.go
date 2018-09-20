package main

import "fmt"

type Person struct{
	name string
}

type Android struct{
	Person 
	model string
}

func (p *Person) talk(){
	fmt.Println("Hi my name is ", p.name)
}

func main(){
	var android *Android = new(Android)

	android.name = "Diego"
	//android.person.talk()
	android.talk()

	/**
	 * To embed method directly then use in type Android the following code
	 * type Android struct{
	 * 	Person
	 * 	model string
	 * }
	 *
	 * And should use directly methods in person-struct
	 *
	 * android.talk() instead of android.person.talk() that is weak
	 */
}