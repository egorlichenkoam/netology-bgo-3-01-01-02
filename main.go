package main

/**
Домашняя работа bgo-3-01-02

https://github.com/netology-code/bgo-homeworks/tree/master/01_std
*/func main() {
	var balance int64 = 1500000000  // 15 миллионов в копейках
	var transfer int64 = 1000000000 // 10 миллионов в копейках
	total := balance + transfer     // int64 + int64 будет int64
	println(total)
}
