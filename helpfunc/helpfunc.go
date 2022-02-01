package helpfunc

import "fmt"

var varFunction = varFun

func varFun() int {
	fmt.Println("varFunc 将被最先执行")
	return 0
}

func Fibonacci(n int) int {
	if n < 2 {
		return 1
	}
	return Fibonacci(n-1) + Fibonacci(n-2)
}

func PrintFibonacci(n int) {
	for i := 0; i < n; i++ {
		fmt.Println(Fibonacci(i))
	}
}

// 先执行变量初始化再执行init()函数最后执行main()
func init() {
	fmt.Println("This is init function,this function will be called before main function")
	fmt.Printf("%d\n", varFunction())
}

// UnNameFunc 匿名函数
func UnNameFunc() {
	// 匿名函数定义的时候直接调用给t赋值
	t := func(a, b int) int {
		return a + b
	}(10, 12)
	fmt.Println("10+12=", t)
}

// ClosureFunc 闭包
func ClosureFunc() func(x int) int {
	var n = 0
	return func(x int) int {
		n = n + x
		return n
	}
}

// DeferDemo key word: defer
// 后入先出的顺序执行。defer 入栈时，会将当时相关的值拷贝
func DeferDemo() int {
	a := 23
	defer fmt.Printf("will print defer1 after return %v\n", a) // a=23
	a = 10
	defer fmt.Printf("will print defer2 after return %v\n", a) // a=10
	fmt.Println("do something")
	a = 20
	return a
}
