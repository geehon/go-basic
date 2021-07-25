package typeint

import "fmt"

func PrintInt() {
	var (
		name string
		age  int
	)
	nickname, height := "gh", 182.1
	name = "GeeHon"
	age = 18

	// TODO:格式化输出
    // %b    表示为二进制
    // %c    该值对应的unicode码值
    // %d    表示为十进制
    // %o    表示为八进制
    // %q    该值对应的单引号括起来的go语法字符字面值，必要时会采用安全的转义表示
    // %x    表示为十六进制，使用a-f
    // %X    表示为十六进制，使用A-F
    // %U    表示为Unicode格式：U+1234，等价于"U+%04X"
    // %E    用科学计数法表示
    // %f    用浮点数表示
	fmt.Printf(name+" is %d years old\n", age)
	fmt.Printf(name+"'s nickname is %s and height is %.2f meter\n", nickname, height)
}
