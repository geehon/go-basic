package array

import "fmt"

func PrintArray() {
	arr := [10]int{2, 3, 4, 5}
	for index, value := range arr {
		fmt.Printf("arr[%d] => %d\n", index, value)
	}
}
