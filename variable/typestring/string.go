package typestring

import "fmt"

func PrintString() {
	var name = "GeeHon"
	greeting := " 学 Golang"
	fmt.Println(name + greeting)

	// 空白标识符 _ 也被用于抛弃值，如值 5 在：_, b = 5, 7 中被抛弃。
	// _ 实际上是一个只写变量，你不能得到它的值。
	_, nickname := getName()
	fmt.Println(nickname)
}

func getName() (userName, nickName string) {
	return "GeeHon", "Gh"
}
