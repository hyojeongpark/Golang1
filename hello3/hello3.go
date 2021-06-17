package main

import "fmt"

var g int = 10 //선언하고 초기화 (가능)
//h:=20 //var h int h=20와 같은 말이지만 /선언후 대입하였기 때문에 (불가능)

func main() {
	var m int = 20
	{
		var s int = 30
		fmt.Println(m, s, g)
	}

}
