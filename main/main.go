package main

import (
	"fmt"
	"github.com/geehon/go-basic/array"
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
}
