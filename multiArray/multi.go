package main

import "fmt"

func main() {
	fmt.Println("다차원 배열 실습")

	//1차원 배열
	var singleArray = [100]int{2: 12, 99: 5} // 두번재 값은 12, 99는 5, 나머지는 0
	fmt.Println(singleArray)
	//2차원 배열
	var doubleArray = [2][3]int{
		{1, 2, 3},
		{2: 55},
	}
	fmt.Println(doubleArray)
	//다차원 배열
	var multiArray = [2][3][2]int{
		{{1: 6}, {5, 7}},
		{
			{0: 3},
			{1: 6},
		},
	}
	fmt.Println(multiArray)
	// 다차원 배열2
	var multiArray2 = [2][3][2]int{
		{
			{1, 2},
			{3},
		},
		{
			{1},
			{1: 2},
		},
	}
	fmt.Println(multiArray2)
}
