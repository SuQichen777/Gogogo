package main

import (
	concurrency "Gogogo/concurrency"
	"fmt"
)

func testChannelSum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum
}
func main(){
	go concurrency.Concurrent_test("Audi")
    go concurrency.Concurrent_test("BMW")
    go concurrency.Concurrent_test("Tesla")
    concurrency.Concurrent_test("Toyota")

	s := []int{1,2,3,4,5,6,7,8,9}
	c := make(chan int)
	go testChannelSum(s[len(s)/2:], c)
	go testChannelSum(s[:len(s)/2], c)
	s1, s2 := <-c, <-c
	fmt.Println("The sum is", s1+s2)
}
