package main

import "fmt"

var k int16 //선언declaration
//k=2 이런건 연산자statement = 함수 외부에 불가
func main() {
	//k = -32768 //int16은 -32768~32767
	//var i int8 //uint는 0~255사이 int는 -128~127사이(0포함이기때문)
	//var j uint8
	//여기만 수정하면됨
	//i = 127
	//j = 255
	//k = 32767
	//fmt.Println("Hello Go World", i, j, k)

	var f float32
	f = 3.14159248
	fmt.Println("실수값 파이 출력: ", f)

	test()
	test1()
}

func test() {
	var a int8
	a = 127
	fmt.Println(a, k)
}