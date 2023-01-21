package pkg1

import "fmt"

// CreateP1 这里是pgk2包下的函数CreateP1
func CreateP1() {
	fmt.Println("这里是pgk1包下的函数CreateP1")
}
func CreateT() {
	fmt.Println("测试下")
}
