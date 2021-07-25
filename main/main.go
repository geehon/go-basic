package main

import (
    "geehon.top/basic/variable/typefloat"
    "geehon.top/basic/variable/typeint"
    "geehon.top/basic/variable/typepointer"
    "geehon.top/basic/variable/typestring"
)

func learnVariable() {
    typestring.PrintString()
    typeint.PrintInt()
    typefloat.PrintFloat()
    typepointer.PrintPointer()
}

func main() {
    learnVariable()
}
