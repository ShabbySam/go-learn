package main

import "go-learn/syntax/variables/demo"

var (
	b = "b"
	c = "c" // 类型推断
) // 块变量（包）（全局）声明

func main() {
	var a int = 1 // 声明局部变量并初始化
	println(a, b, c)

	println(demo.Name) // 大写就是把变量暴露出去
}
