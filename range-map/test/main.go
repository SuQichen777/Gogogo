package main
import (
	"fmt"
	nestmap "Gogogo/range-map"
)

type Vertex struct {
	Lat, Long float64
}

var m = map[string]Vertex{
	"Bell Labs": Vertex{40.1, -74.3},
	"Google": Vertex{437.42202, -122.08408},
}

func updatePrice(prices map[string]float64, fruit string, price float64) {
    prices[fruit] = price
}

func main(){
	nums := []int{1,2,3,4}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println(sum)

	cars := map[string]string{"A": "Audi", "B": "BMW", "T": "Tesla"}
	for k, v := range cars {
		fmt.Printf("%s -> %s\n", k, v)
	}

	fmt.Println(m);

	fruitPrices := map[string]float64{
        "苹果": 5.5,
        "香蕉": 3.8,
    }
	updatePrice(fruitPrices, "苹果", 6.7)
	fmt.Printf("更新后苹果的价格：%.2f\n", fruitPrices["苹果"])
	// 会得到6.7，证明Function传递map是传递的底层指针的拷贝

	nestedRes := nestmap.GetDefualtNestedMap()
	// 获取张三的数学成绩
	if engScore, ok := nestedRes["张三"]["数学"]; ok{
		fmt.Println("张三的数学成绩是", engScore)
	} else {
		fmt.Println("未找到张三的数学成绩")
	}
	fmt.Println(nestedRes)
}
