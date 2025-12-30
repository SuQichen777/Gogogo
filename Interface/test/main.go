package main

import (
	itface "Gogogo/Interface"
	"fmt"
)

func main() {
	fmt.Println("开始main测试")
	itface.StartITFacePkt()
	var c itface.Car

	c = &itface.Tesla{}
	itface.StartCarTest(c)
	c = &itface.Audi{}
	itface.StartCarTest(c)

	itface.TestStringer()
}
