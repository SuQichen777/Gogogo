package errortest

import (
	"errors"
	"fmt"
)

// errors.Join得到的是一个复合error，但是它的return type依然是error
// errors.Join 会自动忽略 nil，并且：
// - 所有参数都是 nil → 返回 nil
// - 部分是 nil → 只 Join 非 nil 的 error
// - 一个非 nil → 返回一个非 nil 的 error
func GetJoinError() error{
	err := errors.Join(nil, nil)
	fmt.Println("正在打印Join两个nil时err是否为nil： ", err == nil) // true

	err1 := NewErrorTest1
	err = errors.Join(nil, err1)
	fmt.Println("正在打印Join一个nil一个非nil： ", err)

	return err
}