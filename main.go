package main

import (
	"fmt"
	hello_world "main/example_1"
	for_loop "main/example_2"
	switches "main/example_3"
	array_slice "main/example_4"
	maps "main/example_5"
	ranger "main/example_6"
)
var funcCounter int = 1
func main() {
	Exercises()
	hello_world.Runner()
	Exercises()
	for_loop.Runner()
	Exercises()
	switches.Runner()
	Exercises()
	array_slice.Runner()
	Exercises()
	maps.Runner()
	Exercises()
	ranger.Runner()
	Exercises()
	Exercises()
	Exercises()

}

func Exercises()  {
	fmt.Printf("\n%d.th Exercise\n",funcCounter)	
	funcCounter ++
}
