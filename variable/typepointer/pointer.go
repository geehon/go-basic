package typepointer

import "fmt"

func getPointer() (int, *int) {
    p := 12
    pointer := &p
    return p, pointer
}

func PrintPointer() {
    _, pointer := getPointer()
    fmt.Printf("指针指向地址的存储的值是:%d", *pointer)
}
