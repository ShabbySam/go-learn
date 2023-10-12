package main

import "fmt"

func main() {
	world := "hello world"
	var a byte = 'a'
	println(world)
	println("%s", world)
	println("hello world" + "hello world") //hello worldhello world
	fmt.Printf("%s", world)
	println(len("n你好")) // 7
	println(a)            // 97
}
