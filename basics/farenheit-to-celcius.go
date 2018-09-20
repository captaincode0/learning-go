package main

import "fmt"

func main(){
  var farenheit_degrees float32
  var celcius_degrees float32

  fmt.Println("Farenheit to celcius conversor")
  fmt.Println("How many farenheit degrees there is? ")
  fmt.Scanf("%f", &farenheit_degrees)

  fmt.Println(farenheit_degrees)

  celcius_degrees = (farenheit_degrees-32.0)*5/9

  fmt.Println("Celcius degrees are ", celcius_degrees)
}
