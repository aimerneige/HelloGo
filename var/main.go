package main

import "fmt"

var msg string = "Author: AimerNeige"
var flag bool
var empty string
var score int

func main() {
	var foo string = "Hello"
	var bar string = "World"
	fmt.Println(foo, bar)

	num := 1
	fmt.Println(num)

	fmt.Println("flag:", flag)
	fmt.Println("empty:", empty)
	fmt.Println("score:", score)

	// var i int, j int // wrong
	// var i, j int
	// fmt.Println(i, j)

	// var i int, j int = 1, 2 // wrong
	var i, j int = 1, 2
	fmt.Println(i, j)

	var m, n = "aa", 1
	fmt.Println(m, n)
}
