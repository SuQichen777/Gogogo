package itface

import "fmt"

type Animal struct{
	name string
	count int
}

func (a Animal) String() string {
	return fmt.Sprintf("%v (count is %v )", a.name, a.count)
}

func TestStringer(){
	cat := Animal{"Cat", 42}
	fmt.Println(cat.String())
}