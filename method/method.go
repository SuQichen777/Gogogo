package method

import "fmt"

type TestType struct{
	Name string
	Age int
}

func StartMethodPkt() {
	fmt.Println("开始Method Packet的测试")
}

func (t TestType) SingleTest(){
	fmt.Println("他的名字是: ", t.Name)
	fmt.Println("他的年龄是: ", t.Age)
	fmt.Println("Single Test内年龄+1")
	t.Age ++
	fmt.Println("他的年龄是: ", t.Age)
}

func (t *TestType) PtrTest(){
	fmt.Println("他的名字是: ", t.Name)
	fmt.Println("他的年龄是: ", t.Age)
	fmt.Println("Ptr Test内年龄+1")
	t.Age ++
	fmt.Println("他的年龄是: ", t.Age)
}