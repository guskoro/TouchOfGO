package main

import "fmt"

var w string

func main() {
	fmt.Println("Hello World !")
	fmt.Println("I want to create new function")

	foo()
	
	fmt.Println("I want to show prime number")

	w = "Ditanyain, ditaruk, dimakan, diminum, ditampar"

	for i:= 0; i<=10; i++ {
		if i%2 == 1 || i == 2 {
			fmt.Println(i)
		}
	}

	boo()
}

func foo ()  {
	fmt.Println("Foo is wakeup after your call")
}

func boo() {
	fmt.Println("Bye lur")
}