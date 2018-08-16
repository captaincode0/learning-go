package main

import "fmt"

func main(){
  var feets float32

  fmt.Println("Feets: ")
  fmt.Scanf("%f", &feets)

  output := feets*0.3048
  fmt.Println("Meters: ", output)
}
