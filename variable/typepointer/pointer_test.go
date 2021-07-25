package typepointer

import "testing"

func TestGetPointer(t *testing.T) {
    if pt, pointer := getPointer(); pt != *pointer && pt == 12 {
        t.Errorf("pt expected be 12, but %d got", pt)
    }
}

func ExamplePrintPointer() {
    PrintPointer()
    // Output: 指针指向地址的存储的值是:12
}
