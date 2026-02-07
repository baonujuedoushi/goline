package main

// 泛型响应结构体：复习你熟悉的广告数据处理
type AdResponse[T any] struct {
	ID    string
	Data  T
	Price float64
}

func fetchAd() {
	// 模拟请求耗时

}

func doAd() {

}
