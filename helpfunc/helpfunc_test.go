package helpfunc

func ExamplePrintFibonacci() {
	PrintFibonacci(10)
	// Output: 1
	//1
	//2
	//3
	//5
	//8
	//13
	//21
	//34
	//55
}

func ExampleDeferDemo() {
	DeferDemo()
	// Output: do something
	//will print defer2 after return 10
	//will print defer1 after return 23
}
