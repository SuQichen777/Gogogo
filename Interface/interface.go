package itface

import "fmt"

type TestType struct{
	Name string
	Age int
}

// 定义一个接口
type Reader interface{
	Read(p []byte) (n int, err error)
}
type Car interface{
	Drive()
}

// 实现一个接口
type File struct {}
func (f *File) Read(p []byte) (int, error) {
	return len(p), nil
}

type Tesla struct {}
type Toyota struct {}
type Audi struct {}

func (t * Tesla) Drive(){
	fmt.Println("I am driving a Tesla.")
}
func (t * Toyota) Drive(){
	fmt.Println("I am driving a Toyota.")
}
func (a * Audi) Drive(){
	fmt.Println("I am driving a Audi.")
}


func StartITFacePkt() {
	fmt.Println("开始Interface Packet的测试")
}

func StartCarTest(c Car) {
	c.Drive()
}

