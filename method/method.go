package method

import (
	"fmt"
	"strings"
)

func Demo() {
	// 实例化结构体
	employee := Employee{id: 1, name: "萧何", age: 50, isLeader: false}
	var employee2 Employee
	employee2.id = 2
	employee2.name = "刘邦"
	employee2.age = 32
	employee2.isLeader = true
	id := employee.getEmployeeId()
	name := employee.getEmployeeName()
	fmt.Printf("员工%v的id：%v\n", name, id)
	fmt.Println("领导的信息:")
	fmt.Printf("id:%v\n姓名:%v\n年龄:%v\n", employee2.getEmployeeId(), employee2.getEmployeeName(), employee2.age)
	employee.setEmployeeName("诸葛亮")
	fmt.Println("指针", employee)
	// setEmployeeID 方法 接收方变量不是指针，修改的不是指针的引用
	employee.setEmployeeID(3)
	fmt.Println("非指针", employee)

	fmt.Println("其他自定义类型的方法")
	testStr := upperString("Hello go")
	fmt.Println(testStr.Upper())
}

type Employee struct {
	id       int
	name     string
	age      int
	isLeader bool
}

func (employee Employee) getEmployeeId() int {
	return employee.id
}

func (employee Employee) getEmployeeName() string {
	return employee.name
}

func (employee Employee) setEmployeeID(id int) {
	employee.id = id
}

func (employee *Employee) setEmployeeName(name string) {
	employee.name = name
}

type upperString string

// Upper go无法实现在不改变地址的情况下改变字符串的内容
func (s upperString) Upper() string {
	return strings.ToUpper(string(s))
}
