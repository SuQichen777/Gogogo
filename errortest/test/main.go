package main

import (
  "Gogogo/errortest"
  "fmt"

)
const (
	errorTestConst1 = 1
	errorTestConst2 = 1
	errorTestConst3 = 1
)

func checkError(testData []int) error {
	for i:=0; i < len(testData); i++{
		if testData[i] == 1 {
			return errortest.NewErrorTest1
		}
	}
	return  nil
}


func main(){
	fmt.Println("基础 测试")
	testData1 := [3]int{errorTestConst1, errorTestConst2, errorTestConst3}
	err := checkError(testData1[:])
	fmt.Println(err)
	fmt.Println("---------------------------------------")
	fmt.Println("Join 测试")
	errortest.GetJoinError()
}