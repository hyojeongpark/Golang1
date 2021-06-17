package main

//import "fmt"
//import "unsafe" 저장하면 없어짐
import (
	"fmt"
	"unsafe"
)

func main() {
	a := 1                   //int (int64라고 안하는건 체제에 맞게 유연하게 하려고)
	var b = 3.14             //float
	c := "Hello world"       //string
	d := int32(10)           //int32
	var e float32 = 3.141592 //float32

	var f int16 = 256 //2의 8승
	var g int8 = int8(f)
	var h int8 = -128
	var s string = "print S"

	fmt.Println("Hi", a, b, c, d, e)
	fmt.Println(g)
	fmt.Println("이진수: %b, 십진수 : %d", h, h)
	fmt.Printf("test %s\n", s)

	var q int
	var r bool = true
	fmt.Println(r)
	fmt.Printf("양식으로 출력 %t, %T\n", r, q)
	fmt.Println(unsafe.Sizeof(q)) //숫자로 호출(윈도우체제)
	fmt.Println(unsafe.Sizeof(r)) //전체 8 바이트이기 때문에 7바이트를 낭비하는 boolean...
	fmt.Println(StringReturn())
}

func StringReturn() string {
	return "ABC"

}
