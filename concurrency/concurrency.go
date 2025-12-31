package concurrency

import (
	"fmt"
	"time"
)

func Concurrent_test(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}


