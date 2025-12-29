package nestmap

var TestInt bool = true

func GetDefualtNestedMap() map[string]map[string]int{
	res := make(map[string]map[string]int)
	res ["张三"] = map[string]int{
		"语文": 100,
		"数学": 90,
	}

	return res
} 
