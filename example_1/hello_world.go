package hello_world

import ("fmt")
func Runner() {
	a:= "HELLO WORLD"
	b:= "This is the first example of the code"
	c:="-----"
	e:="new"
	d:= fmt.Sprintf("%s sentence\t",e)
	fmt.Printf("%s\t%s %s\n%s. ",c,a,c,b)
	println(d,"new","string")
	fmt.Print(d,"new","string","in","here","too")
	println(d)
}