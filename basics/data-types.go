package main

import "fmt"

func main(){
  /**
    Integers
    Are numbers without a decimal representation
    In go there are different kind of integers
    1. uint: unsigned integer with domain from [0,n]
    2. int: signed integer with domain from [-n,n]

    There are three machine integer that depends of size that architecture defines
    uint, int, uintptr
  */
  var t uint8

  t=(100+20)
  fmt.Println(t)

  /**
    Float point numbers
    This kind of numbers are also known as Real Numbers
    and contains a decimal representation 
    There are two kind of float point numbers
    1. float32: float with 32 bits for decimal representation.
    2. float64: float with 64 bits for decimal representation.

    Also exists one implementation for complex numbers
    1. complex64
    2. complex128
  */

  var r float32

  r=1.00000032+1.000000064265;
  fmt.Println(r)

  /**
    String
    One string is a set of characters like [hello]

    len: determines the size of one string
  */
  var s string

  s = "some string"

  fmt.Println(len(s))
  fmt.Println(s)
  fmt.Println(s[0]) //get one character int representation with numerical index

  /**
    Boolean datatype
    This kind of datatype has two different values [false,true]
    is used for logical operations
 
    Logic operators:
    1. &&: and operator
    2. ||: or operator
    3. !: not operator
    
  */
  var w bool = false

  fmt.Println(w)
}
