package main

import (
	method "Gogogo/method"
	"fmt"
)

func main() {
	fmt.Println("开始main测试")
	method.StartMethodPkt()
	testPerson := method.TestType{Name: "张三", Age: 30}
	testPerson.SingleTest()
	fmt.Println("再次进行single test")
	testPerson.SingleTest()
	fmt.Println("进行pointer test")
	testPerson.PtrTest()
	fmt.Println("再次进行pointer test")
	testPerson.PtrTest()
}
