package main

import "fmt"

func main() {
	a := 3 //a변수는 운영체제(32or64비트)의 사이즈를 따라서 결정된다
	var b float64 = 3.4

	//var c int = b //실수에 정수를 넣는 것은 불가능
	var c int = int(b) //실수가 정수로 바뀜
	//var c float64 = b
	//d := a * b //a*b의 타입이 정해져 있지 않기 때문에 에러
	d := float64(a) * b
	fmt.Println("Print", c, d)
}
