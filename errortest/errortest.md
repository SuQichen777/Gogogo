# errortest
## 简介
对于`Go`语言`errors`库的测试与使用
## 已完成
- `errors.New`: 详见 [error.go](https://github.com/SuQichen777/Gogogo/blob/main/errortest/error.go)
    - 官方定义：`func New(text string) error`。创建一个固定消息的 error。
- `errors.Join`: 详见 [joinerror.go](https://github.com/SuQichen777/Gogogo/blob/main/errortest/joinerror.go)
    - 官方定义：`func Join(errs ...error) error`接受 0 个、1 个或多个 error, 也可以用`...`语法糖直接传入一个error slice：
    ``` Go
    var errs []error

    errs = append(errs, err1)
    errs = append(errs, err2)

    return errors.Join(errs...)
    ```
## TODO
- `errors.Is`
    - 官方定义：`func Is(err, target error) bool`。err 这个 error（包括 wrap / join 里的）是不是 包含 target。
    - 应该使用`errors.Is`进行比较，而不是`==`
- `errors.As`
    - 官方定义：`func As(err error, target any) bool`。As 常用于提取结构化信息（比如路径、超时、状态码等）。