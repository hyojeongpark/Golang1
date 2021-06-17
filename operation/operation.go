package main

import (
	"fmt"
)

func main() {
	//var a int8
	//a = 112
	//var b int8
	//b = 110
	//c := a & b
	//d := a | b
	//e := a ^ b
	//f := a &^ b
	//var g = 4
	//h := g << 2
	i := 1
	j := 1
	var p1 *int = &i
	var p2 *int = &j
	fmt.Println(p1 == p2)
	fmt.Printf("p1과 p2의 주소는: %p 와 %p\n", p1, p2)
	fmt.Printf("가리키는 값은 %d 와 %d\n", *p1, *p2)

	var x int16 = 7
	y := ^x          // 반전시킨것과
	var z int16 = -7 //그냥 (-)음수를 만드는 차이점

	fmt.Printf("%016b %d %016b\n", y, y, z)

	//fmt.Printf("%08b\n", c)
	//fmt.Printf("%8b\n", d)
	//fmt.Printf("%08b\n", e)
	//fmt.Printf("%08b\n", f)
	//fmt.Printf("%08b\n", h)
	//fmt.Printf("ABC %b DEF \n", 13)
	//fmt.Printf("%c%c \n", 'g', 'o')
	//fmt.Printf("%04b %d \n", h, h>>1) //4는 이진수로100, 10000에서 오른쪽으로 한칸하면 1000(이진수로8)

}
