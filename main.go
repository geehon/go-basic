package main

import (
	"fmt"
	"github.com/geehon/go-basic/array"
	"github.com/geehon/go-basic/helpfunc"
	"github.com/geehon/go-basic/method"
	"github.com/geehon/go-basic/variable/typefloat"
	"github.com/geehon/go-basic/variable/typeint"
	"github.com/geehon/go-basic/variable/typepointer"
	"github.com/geehon/go-basic/variable/typestring"
)

func learnVariable() {
	typestring.PrintString()
	typeint.PrintInt()
	typefloat.PrintFloat()
	typepointer.PrintPointer()
}

func createLine(str string) {
	fmt.Printf("=================== %s ===================\n", str)
}

func main() {
	learnVariable()
	createLine("print array")
	array.PrintArray()
	createLine("print fibonacci")
	helpfunc.PrintFibonacci(10)
	createLine("匿名函数")
	helpfunc.UnNameFunc()
	createLine("闭包")
	f := helpfunc.ClosureFunc()
	f(1)
	f(1)
	f(2)
	fmt.Println("调用4次后计数器值：", f(1))
	createLine("defer demo")
	helpfunc.DeferDemo()
	createLine("builtin string function")
	helpfunc.StringBuiltin()
	createLine("use method in go")
	method.Demo()
}
