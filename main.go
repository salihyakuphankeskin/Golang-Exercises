package main

import (
	"fmt"
	hello_world "main/example_1"
	for_loop "main/example_2"
)
var funcCounter int = 1
func main() {
	Exercises()
	hello_world.Runner()
	Exercises()
	for_loop.Runner()
	Exercises()
	Exercises()
	Exercises()
	Exercises()
	Exercises()
	Exercises()
	Exercises()

}

func Exercises()  {
	fmt.Printf("\n%d.th Exercise\n",funcCounter)	
	funcCounter ++
}
