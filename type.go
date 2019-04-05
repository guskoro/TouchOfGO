package main

import "fmt"

var a int
type huruf string
var b huruf

func main()  {
	a = 9
	b = "guskoro"
	fmt.Println(a)
	fmt.Printf("%T\n", a)
	fmt.Println(b)
	fmt.Printf("%T\n", b)
}