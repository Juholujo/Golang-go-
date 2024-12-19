package main

type T [2]interface{}

func main() {
	var a = T{1, func() {}}
	var b = T{2, func() {}}
	println(a == b) // false

	//var c = T{2, func(){}}
	//var d = T{2, func(){}}
	//println(c == d) // panics

	var x = T{2, "Andrew"}
	var x2 = T{2, "Andrew"}
	println(x == x2)

}
